package httpservice

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/ManiMuridi/goclean/util"

	"github.com/ManiMuridi/goclean/service"

	"github.com/ManiMuridi/goclean/config"

	"github.com/ManiMuridi/goclean/validation"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type HttpService interface {
	Run()
	Handler() HttpHandler
	Validator() *validation.V
}

type httpService struct {
	handler   HttpHandler
	http      *echo.Echo
	validator *validation.V
	logger    *zerolog.Logger
}

func (s *httpService) ConfigLogger(logger *zerolog.Logger) {
	s.logger = logger
}

func NewHttp(handler HttpHandler) HttpService {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	e := echo.New()

	e.Use( // middleware will get the preferred language from the Accept-Langauge header value
		func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				if lang := c.Request().Header.Get("Accept-Language"); lang != "" {
					util.Tr.SetLanguage(lang)
				}
				return next(c)
			}
		})

	svc := service.Bootstrap(&httpService{
		handler: handler,
		http:    e,
	})

	return svc.(HttpService)
}

func (s *httpService) Validator() *validation.V {
	return s.validator
}

func (s *httpService) configure() error {
	s.validator = validation.Validator
	s.http.Validator = s.validator

	return nil
}

func (s *httpService) Bootstrap() error {
	if err := s.configure(); err != nil {
		return err
	}

	s.logger.Debug().Msg("Bootstrapping " + s.Name() + " Service")

	for i := range s.handler.Routes() {
		route := s.handler.Routes()[i]

		handlerFunc := func(ctx echo.Context) error {
			return route.Handler(&Context{Context: ctx})
		}

		s.http.Add(route.Method, route.Path, handlerFunc, route.Middleware...)
	}

	for i := range s.handler.Middleware() {
		s.http.Use(s.handler.Middleware()[i])
	}

	return nil
}

func (s *httpService) Name() string {
	return config.GetString("service.name")
}

func (s *httpService) Handler() HttpHandler {
	return s.handler
}

func (s *httpService) Run() {
	go func() {
		s.logger.Info().Msg("Running service...")

		if err := s.http.Start(fmt.Sprintf(":%d", config.GetInt("http.port"))); err != nil {
			s.http.Logger.Info("shutting down the server")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.http.Shutdown(ctx); err != nil {
		s.http.Logger.Fatal(err)
	}
}

func (s *httpService) Logger() *zerolog.Logger {
	return s.logger
}

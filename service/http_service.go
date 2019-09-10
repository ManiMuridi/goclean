package service

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/ManiMuridi/goclean/config"

	"github.com/ManiMuridi/goclean/translator"

	"github.com/ManiMuridi/goclean/validation"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/go-playground/validator.v9"
)

type HttpService interface {
	Run()
	Handler() HttpHandler
	//Config() *HttpConfig
	Validator() *validation.Validator
}

//type HttpConfig struct {
//	Name string
//	Port uint16
//}

type httpService struct {
	handler HttpHandler
	http    *echo.Echo
	//config    *HttpConfig
	validator *validation.Validator
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
				lang := c.Request().Header.Get("Accept-Language")
				translator.Tr.SetLanguage(lang)
				return next(c)
			}
		})
	svc := Bootstrap(&httpService{
		//config:  config,
		handler: handler,
		http:    e,
	})

	return svc.(HttpService)
}

//func (s *httpService) Config() *HttpConfig {
//	return s.config
//}

func (s *httpService) Validator() *validation.Validator {
	return s.validator
}

func (s *httpService) configure() error {
	validate := validator.New()

	s.validator = &validation.Validator{Validator: validate}
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
		s.http.Add(route.Method, route.Path, route.HandlerFunc, route.Middleware...)
	}

	for i := range s.handler.Middleware() {
		s.http.Use(s.handler.Middleware()[i])
	}

	return nil
}

func (s *httpService) Name() string {
	return config.Service.Name
}

func (s *httpService) Handler() HttpHandler {
	return s.handler
}

func (s *httpService) Run() {
	go func() {
		s.logger.Info().Msg("Running service...")

		if err := s.http.Start(fmt.Sprintf(":%d", config.Http.Port)); err != nil {
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

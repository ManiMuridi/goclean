package service

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// TODO add domain events interface and service
type Service interface {
	Name() string
	Bootstrap() error
	Run()
	ConfigLogger(logger *zerolog.Logger)
	Logger() *zerolog.Logger
}

func Bootstrap(svc Service) Service {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	logger := log.With().
		Str("service", svc.Name()).
		Logger()

	svc.ConfigLogger(&logger)

	if len(strings.TrimSpace(svc.Name())) == 0 {
		log.Panic().Msg("service name is empty")
		return nil
	}

	if err := svc.Bootstrap(); err != nil {
		panic(err)
	}

	return svc
}

package brokerservice

import (
	"os"

	"github.com/ManiMuridi/goclean/service"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
)

var (
	persistanceEnabled = false
	storage            Storage
)

type Handler interface {
	Queues() []Queue
}

type QueueHandlerFunc func(delivery amqp.Delivery)

type Queue struct {
	Exchange    string
	Topic       string
	queue       amqp.Queue
	HandlerFunc QueueHandlerFunc
}

type Exchange struct {
	Name    string
	channel *amqp.Channel
	Queues  []Queue
}

func SetStorage(store Storage) {
	storage = store
	persistanceEnabled = true
}

func NewConsumer(handler Handler) Consumer {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	svc := service.Bootstrap(&consumer{
		handler:   handler,
		exchanges: make(map[string]Exchange),
	})

	return svc.(Consumer)
}

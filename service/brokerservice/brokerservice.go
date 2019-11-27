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

func declareNamedQueue(ch *amqp.Channel, queueName string) (amqp.Queue, error) {
	return ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
}

func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
}

func declareExchange(ch *amqp.Channel, exchangeName string) error {
	return ch.ExchangeDeclare(
		exchangeName, // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
}

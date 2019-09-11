package brokerservice

import (
	"os"

	"github.com/ManiMuridi/goclean/service"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
)

/*
TODO add storage interface
TODO add domain events interface and service
Message
Session
Connection
Channel
Queue


Broker (Server): An application - implementing the AMQP model - that accepts connections from clients for message routing, queuing etc.
Message: Content of data transferred / routed including information such as payload and message attributes.
Consumer: An application which receives message(s) - put by a producer - from queues.
Producer: An application which put messages to a queue via an exchange.


Exchange: A part of the broker (i.e. server) which receives messages and routes them to queues
Queue (message queue): A named entity which messages are associated with and from where consumers receive them
Bindings: Rules for distributing messages from exchanges to queues
*/

//consumer := amqp.NewConsumer(&config{
//Exchange: "topic-exchange",
//Url:      "amqp://rabbitmq:rabbitmq@rabbitmq.quantum.stp:30232",
//})
//
//consumer.HandleTopic("routing.key.or.topic", func(delivery amqp.Delivery) {
//	ProcessResult(delivery)
//})
//
//consumer.HandleTopic("routing.key.or.topic.two", func(delivery amqp.Delivery) {
//	ProcessResult(delivery)
//})
//
//consumer.Run()

// type Exchange interface{}

// type Queue interface{}

// type Binding interface{}

// type Broker interface {
// 	Connection()
// 	Channel()
// }

// type topicHandlerFunc func(d amqp.Delivery)

// type BrokerConfig struct {
// 	Url      string
// 	Exchange string
// }

// type Producer interface{}

// type Message interface{}
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

//type Config struct {
//	ServiceName string
//	Exchange    string
//	Topic       string
//	Url         string
//}

func SetStorage(store Storage) {
	storage = store
	persistanceEnabled = true
}

func NewConsumer(handler Handler) Consumer {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	svc := service.Bootstrap(&consumer{
		//config:    config.Service,
		handler:   handler,
		exchanges: make(map[string]Exchange),
	})

	return svc.(Consumer)
}

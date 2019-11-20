package brokerservice

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/ManiMuridi/goclean/config"

	"github.com/ManiMuridi/goclean/validation"
	"github.com/rs/zerolog"
	"github.com/streadway/amqp"
)

type Consumer interface {
	Run()
	Handler() Handler
	//Config() *Config
	Validator() *validation.V
}

type consumer struct {
	handler    Handler
	exchanges  map[string]Exchange
	connection *amqp.Connection
	validator  *validation.V
	//config     *Config
	logger *zerolog.Logger
}

func (c *consumer) ConfigLogger(logger *zerolog.Logger) {
	c.logger = logger
}

func (c *consumer) Validator() *validation.V {
	return c.validator
}

func (c *consumer) Handler() Handler {
	return c.handler
}

func (c *consumer) Name() string {
	return config.GetString("service.name")
}

//func (c *consumer) Config() *Config {
//	return c.config
//}

func (c *consumer) Bootstrap() error {
	if err := c.configure(); err != nil {
		return err
	}

	c.logger.Debug().Msg("Bootstrapping " + c.Name() + " Service")

	c.openConnection()

	for i := range c.handler.Queues() {
		c.logger.Info().Msg("Bootstrapping Exchange..." + c.handler.Queues()[i].Exchange)
		ch, err := c.connection.Channel()

		if err != nil {
			c.logger.Panic().Err(err)
		}

		listener := Queue{
			Exchange:    c.handler.Queues()[i].Exchange,
			Topic:       c.handler.Queues()[i].Topic,
			HandlerFunc: c.handler.Queues()[i].HandlerFunc,
		}
		exchange := Exchange{
			Name:    c.handler.Queues()[i].Exchange,
			channel: ch,
		}

		queue, err := exchange.channel.QueueDeclare(
			"",    // name
			false, // durable queue or temporary
			false, // delete when unused
			true,  // exclusive
			false, // no-wait
			nil,   // arguments
		)

		if err != nil {
			panic(err)
		}

		listener.queue = queue

		err = exchange.channel.QueueBind(
			listener.queue.Name,
			listener.Topic,
			exchange.Name,
			false,
			nil,
		)

		if len(exchange.Queues) == 0 {
			exchange.Queues = make([]Queue, 0)
		}

		exchange.Queues = append(exchange.Queues, listener)
		c.exchanges[exchange.Name] = exchange

	}

	return nil
}

func (c *consumer) Run() {
	defer c.connection.Close()
	var wg sync.WaitGroup
	wg.Add(len(c.exchanges))

	for exchange := range c.exchanges {
		go c.setupConsumer(&wg, exchange)
	}

	wg.Wait()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}

func (c *consumer) setupConsumer(wg *sync.WaitGroup, exchange string) {
	defer wg.Done()

	messages, err := c.exchanges[exchange].channel.Consume(
		"",
		"",
		true,
		false,
		false,
		false,
		nil)

	if err != nil {
		panic(err)
	}

	for d := range messages {
		c.handleMessages(exchange, d)
	}
}

func (c *consumer) handleMessages(exchange string, d amqp.Delivery) {
	for j := range c.exchanges[exchange].Queues {
		if d.RoutingKey == c.exchanges[exchange].Queues[j].Topic {
			if persistanceEnabled {
				var m *Message
				if err := json.Unmarshal(d.Body, &m); err != nil {
					// TODO handle error gracefully
					fmt.Println(err)
				}

				if _, err := storage.Store(m); err != nil {
					// TODO handle error gracefully
					fmt.Println(err)
				}
			}
			c.exchanges[exchange].Queues[j].HandlerFunc(d)
		}
	}
}

func (c *consumer) Logger() *zerolog.Logger {
	return c.logger
}

func (c *consumer) configure() error {
	c.validator = validation.Validator

	return nil
}

func (c *consumer) openConnection() {
	conn, err := amqp.Dial(config.GetString("broker.url"))

	if err != nil {
		c.logger.Panic().Err(err)
	}

	c.connection = conn
}

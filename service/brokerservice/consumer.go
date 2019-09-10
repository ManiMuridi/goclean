package brokerservice

import (
	"os"
	"os/signal"
	"sync"

	"github.com/ManiMuridi/goclean/config"

	"github.com/ManiMuridi/goclean/validation"
	"github.com/rs/zerolog"
	"github.com/streadway/amqp"
	"gopkg.in/go-playground/validator.v9"
)

type Consumer interface {
	Run()
	Handler() Handler
	//Config() *Config
	Validator() *validation.Validator
}

type consumer struct {
	handler    Handler
	exchanges  map[string]Exchange
	connection *amqp.Connection
	validator  *validation.Validator
	//config     *Config
	logger *zerolog.Logger
}

func (c *consumer) ConfigLogger(logger *zerolog.Logger) {
	c.logger = logger
}

func (c *consumer) Validator() *validation.Validator {
	return c.validator
}

func (c *consumer) Handler() Handler {
	return c.handler
}

func (c *consumer) Name() string {
	return config.Service.Name
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

	for k, _ := range c.exchanges {
		go func(k string) {
			defer wg.Done()
			messages, err := c.exchanges[k].channel.Consume("", "", true, false, false, false, nil)

			if err != nil {
				panic(err)
			}

			func(delivery <-chan amqp.Delivery) {
				for d := range delivery {
					for j := range c.exchanges[k].Queues {
						if d.RoutingKey == c.exchanges[k].Queues[j].Topic {
							c.exchanges[k].Queues[j].HandlerFunc(d)
						}
					}
				}
			}(messages)
		}(k)
	}

	wg.Wait()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}

func (c *consumer) Logger() *zerolog.Logger {
	return c.logger
}

func (c *consumer) configure() error {
	validate := validator.New()
	c.validator = &validation.Validator{Validator: validate}

	return nil
}

func (c *consumer) openConnection() {
	conn, err := amqp.Dial(config.Broker.Url)

	if err != nil {
		c.logger.Panic().Err(err)
	}

	c.connection = conn
}

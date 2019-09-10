package main

import (
	"fmt"
	"log"

	"github.com/ManiMuridi/goclean/rabbitmq"

	"github.com/streadway/amqp"
)

const (
	amqpUrl = "amqp://rabbitmq:rabbitmq@rabbitmq.quantum.stp:30232"
)

func main() {
	conn, err := amqp.Dial(amqpUrl)
	defer conn.Close()

	if err != nil {
		panic(err)
	}

	consumer := rabbitmq.NewConsumer(conn, "testing-topic-exchange")

	err = consumer.Listen("", []string{"testing.topic"}, func(delivery <-chan amqp.Delivery) {
		for d := range delivery {
			fmt.Println(fmt.Sprintf("%+v", d.RoutingKey))
			log.Printf(" [x] %s", d.Body)
		}
	})

	if err != nil {
		panic(err)
	}

}

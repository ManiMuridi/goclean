package main

import (
	"fmt"
	"path/filepath"

	"github.com/ManiMuridi/goclean/config"

	"github.com/ManiMuridi/goclean/service/brokerservice"
	"github.com/streadway/amqp"
)

type qHandler struct{}

var (
	exchange1 = "testing-topic-exchange"
	exchange2 = "testing-topic-exchange-create"
	topic     = "testing.created"
)

func (qh *qHandler) Queues() []brokerservice.Queue {
	return []brokerservice.Queue{
		{
			Exchange: exchange1,
			Topic:    topic,
			HandlerFunc: func(d amqp.Delivery) {
				fmt.Println(d.UserId)
				fmt.Println(string(d.Body))
			},
		},
		{
			Exchange: exchange2,
			Topic:    topic,
			HandlerFunc: func(d amqp.Delivery) {
				/* TODO detect if it has been acked or received by a consumer properly in streadway docs */
				fmt.Println(string(d.Body))
			},
		},
	}
}

func main() {
	path, _ := filepath.Abs("../config.toml")

	if err := config.Load(path); err != nil {
		panic(err)
	}

	qlSvc := brokerservice.NewConsumer(&qHandler{})
	qlSvc.Run()
}

package brokerservice

import (
	"encoding/json"
	"time"

	"github.com/ManiMuridi/goclean/config"

	"github.com/google/uuid"

	"github.com/streadway/amqp"
)

var (
	contentType = "application/json"
)

type messageMeta struct {
	ID          uuid.UUID
	ContentType string
	PushedAt    time.Time
}

type Message struct {
	Meta     messageMeta
	Exchange string
	Topic    string
	Payload  interface{}
}

func (m *Message) Bytes() []byte {
	jBytes, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	return jBytes
}

func SetContentType(messageContentType string) {
	contentType = messageContentType
}

func Push(exchange, topic string, payload interface{}) (*Message, error) {
	meta := messageMeta{
		ID:          uuid.New(),
		ContentType: contentType,
		PushedAt:    time.Now(),
	}

	m := Message{
		Meta:     meta,
		Exchange: exchange,
		Topic:    topic,
		Payload:  payload,
	}

	if err := pushMessage(m); err != nil {
		return nil, err
	}

	return &m, nil
}

func pushMessage(msg Message) error {
	conn, err := amqp.Dial(config.Broker.Url)

	if err != nil {
		return err
	}

	channel, err := conn.Channel()

	if err != nil {
		return err
	}

	defer conn.Close()
	defer channel.Close()

	err = channel.Publish(
		msg.Exchange,
		msg.Topic,
		false,
		false,
		amqp.Publishing{
			ContentType: msg.Meta.ContentType,
			Body:        msg.Bytes(),
		},
	)

	if err != nil {
		return err
	}

	return nil
}

//func GenPath(evt interface{}) string {
//	return strings.ToLower(
//		strings.Join(
//			camelcase.Split(
//				reflect.TypeOf(evt).Elem().Name()),
//			"."))
//}

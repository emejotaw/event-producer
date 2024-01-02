package rabbitmq

import (
	"context"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	channel    *amqp.Channel
	exchange   string
	routingKey string
	queue      string
}

func NewRabbitMQ(user, password, host string, port int) (*RabbitMQ, error) {

	rabbitMQ := &RabbitMQ{
		exchange: "amq.direct",
		queue:    "event-queue",
	}
	err := rabbitMQ.connect(user, password, host, port)

	if err != nil {
		return nil, err
	}

	return rabbitMQ, nil
}

func (r *RabbitMQ) connect(user, password, host string, port int) error {

	dsn := fmt.Sprintf("amqp://%s:%s@%s:%d/", user, password, host, port)
	conn, err := amqp.Dial(dsn)

	if err != nil {
		log.Printf("could not establish amqp connection, error: %v", err)
		return err
	}

	channel, err := conn.Channel()

	if err != nil {
		log.Printf("could not connect with channel, error: %v", err)
		return err
	}

	r.channel = channel
	return nil
}

func (r *RabbitMQ) Publish(body []byte) error {

	ctx := context.Background()
	return r.channel.PublishWithContext(ctx,
		r.exchange,
		r.routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

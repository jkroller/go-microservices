package events

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn      *amqp.Connection
	queueName string
}

func NewConsumer(conn *amqp.Connection) (Consumer, error) {
	consumer := Consumer{
		conn: conn,
	}
}

func (consumer *Consumer) setup() error {
	chanel, err := consumer.conn.Channel()
	if err != nil {
		return err
	}

	return
}

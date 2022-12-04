package consumer

import "github.com/streadway/amqp"

type Consumer struct {
	Conn    *amqp.Connection
	channel *amqp.Channel
}

func NewConsumer(address string) (Consumer, error) {
	conn, err := amqp.Dial(address)
	if err != nil {
		return Consumer{}, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return Consumer{}, err
	}

	return Consumer{
		Conn:    conn,
		channel: ch,
	}, nil
}

func (c *Consumer) ConsumeMessage(queueName string) (<-chan amqp.Delivery, error) {
	msgs, err := c.channel.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		return nil, err
	}
	return msgs, nil
}

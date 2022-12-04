package producer

import (
	"github.com/streadway/amqp"
)

type Producer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	Queues  map[string]amqp.Queue
}

func NewProducer(address string) (Producer, error) {

	conn, err := amqp.Dial(address)
	if err != nil {
		return Producer{}, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return Producer{}, err
	}

	return Producer{
		conn:    conn,
		channel: ch,
		Queues:  make(map[string]amqp.Queue),
	}, err
}

func (p *Producer) CreateQueue(nameQueue string) error {
	q, err := p.channel.QueueDeclare(
		nameQueue,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	p.Queues[nameQueue] = q
	return nil
}

func (p *Producer) SendMessage(nameQueue, body string) error {
	err := p.channel.Publish(
		"",
		nameQueue,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		return err
	}
	return nil
}

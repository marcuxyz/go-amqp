package messages

import (
	"log"

	"github.com/streadway/amqp"
)

type Publisher struct {
	ch *amqp.Channel
}

func (p *Publisher) Init() {
	connectRabbitMQ, err := amqp.Dial(ampqServerURL())
	if err != nil {
		log.Fatal(err)
	}
	// defer connectRabbitMQ.Close()

	p.ch, err = connectRabbitMQ.Channel()
	if err != nil {
		log.Fatal(err)
	}
	// defer p.ch.Close()

	_, err = p.ch.QueueDeclare(
		"messages", // queue name
		true,       // durable
		false,      // auto delete
		false,      // exclusive
		false,      // no wait
		nil,        // arguments
	)
	if err != nil {
		log.Fatal(err)
	}
}

func (p *Publisher) Publish() {

	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("msg"),
	}

	// Attempt to publish a message to the queue.
	if err := p.ch.Publish(
		"correios", // exchange
		"messages", // routing_key
		false,
		false,   // immediate
		message, // message to publish
	); err != nil {
		log.Fatal(err)
	}
}

func ampqServerURL() string {
	return "amqp://user:user@localhost:5672"
}

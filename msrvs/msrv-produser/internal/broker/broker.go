package broker

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log/slog"
)

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func (b *RabbitMQ) Send(ctx *context.Context, topic string, body []byte) error {
	q, err := b.ch.QueueDeclare(
		topic, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = b.ch.PublishWithContext(*ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	failOnError(err, "Failed to publish a message")
	slog.Info(" [x] Sent %s\n", string(body))
	return nil
}

func (b *RabbitMQ) Close() {
	failOnError(b.conn.Close(), "Failed to close connection broker")
	failOnError(b.ch.Close(), "Failed to close channel broker")
}

func failOnError(err error, msg string) {
	if err != nil {
		slog.Error(msg, err)
	}
}

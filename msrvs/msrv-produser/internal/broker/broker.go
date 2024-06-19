package broker

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"log/slog"
	"time"
)

type IBroker interface {
	Send(ctx *context.Context, topic string, body string) error
	Close()
}

func InitBroker(addr string) IBroker {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
	for {
		select {
		case <-ctx.Done():
			log.Fatalf("Failed to connect to RabbitMQ")
			return nil
		default:
			conn, err := amqp.Dial(addr)
			if err == nil {
				slog.Info("RabbitMQ connected")
				ch, err1 := conn.Channel()
				failOnError(err1, "Failed to open a channel")
				cancel()

				return &RabbitMQ{conn: conn, ch: ch}
			}

			slog.Error("Failed to connect to RabbitMQ", err)
			time.Sleep(2 * time.Second)
		}
	}
}

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func (b *RabbitMQ) Send(ctx *context.Context, topic string, body string) error {
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
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	slog.Info(" [x] Sent %s\n", body)
	return nil
}

func (b *RabbitMQ) Close() {
	b.conn.Close()
	b.ch.Close()
}

func failOnError(err error, msg string) {
	if err != nil {
		slog.Error(msg, err)
	}
}

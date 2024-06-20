package broker

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"log/slog"
	"time"
)

type IBroker interface {
	Send(ctx *context.Context, topic string, body []byte) error
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

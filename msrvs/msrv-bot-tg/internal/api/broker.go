package api

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"log/slog"
	"projectX/msrvs/msrv-bot-tg/internal/service"
	"time"
)

type IBroker interface {
	WatchEvents()
	Close()
}

func InitBroker(addr string, service service.IServiceSet) IBroker {
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

				return &RabbitMQ{conn: conn, ch: ch, srv: service}
			}

			slog.Error("Failed to connect to RabbitMQ", err)
			time.Sleep(2 * time.Second)
		}
	}
}

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	srv  service.IServiceSet
}

func (r *RabbitMQ) WatchEvents() {
	q, err := r.ch.QueueDeclare(
		"event", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	msgs, err := r.ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			r.srv.Set(context.Background(), d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages")
	<-forever
}

func (r *RabbitMQ) Close() {
	failOnError(r.conn.Close(), "Failed to close connection broker")
	failOnError(r.ch.Close(), "Failed to close channel broker")
}

func failOnError(err error, msg string) {
	if err != nil {
		slog.Error(msg, err)
	}
}

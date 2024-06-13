package api

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"log/slog"
	"os"
	"projectX/msrvs/msrv-bot-tg/internal/service"
)

type IBroker interface {
	Get()
	Close()
}

func InitBroker(addr string, service service.IService) IBroker {
	conn, err := amqp.Dial(addr)
	if err != nil {
		slog.Error("Failed to connect to RabbitMQ", err)
		os.Exit(1)
	}

	ch, err1 := conn.Channel()
	failOnError(err1, "Failed to open a channel")

	return &RabbitMQ{conn: conn, ch: ch, srv: service}
}

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	srv  service.IService
}

func (r *RabbitMQ) Get() {
	q, err := r.ch.QueueDeclare(
		"test", // name
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
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
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func (r *RabbitMQ) Close() {
	r.conn.Close()
	r.ch.Close()
}

func failOnError(err error, msg string) {
	if err != nil {
		slog.Error(msg, err)
	}
}

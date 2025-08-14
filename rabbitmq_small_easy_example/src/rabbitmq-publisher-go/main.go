package main

import (
	"log"
	"os"
	"time"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	rabbitHost := os.Getenv("RABBITMQ_HOST")
	if rabbitHost == "" {
		rabbitHost = "rabbitmq"
	}

	connStr := "amqp://guest:guest@" + rabbitHost + ":5672/"

	var conn *amqp.Connection
	var err error
	for i := 0; i < 5; i++ {
		conn, err = amqp.Dial(connStr)
		if err == nil {
			break
		}
		log.Printf("Connection attempt %d failed: %v", i+1, err)
		time.Sleep(5 * time.Second)
	}
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Declarar una cola (si no existe)
	q, err := ch.QueueDeclare(
		"hello", // nombre de la cola
		false,   // durable (sobrevive a reinicios del servidor)
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// Mensaje a enviar (puede ser un argumento del CLI)
	body := "Hello RabbitMQ!"
	if len(os.Args) > 1 {
		body = os.Args[1]
	}

	// Publicar el mensaje
	for i := 0; i < 10; i++ { // Envía 10 mensajes de ejemplo
		err = ch.Publish(
			"",     // exchange (vacío para el default)
			q.Name, // routing key (nombre de la cola)
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
				Timestamp:   time.Now(),
			},
		)
		failOnError(err, "Failed to publish a message")
		log.Printf(" [x] Sent %s", body)
		time.Sleep(3 * time.Second)
	}
}

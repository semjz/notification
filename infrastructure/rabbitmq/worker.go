package rabbitmq

import (
	"encoding/json"
	"log"
	"notification/pkg/setup"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Receiver() {
	// Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Declare a queue
	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	// Consume messages from the queue
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	// This channel will keep the main goroutine alive
	var forever chan struct{}

	type TypeHint struct {
		Type string `json:"type"`
	}

	// Start a goroutine to process incoming messages
	go func() {
		for d := range msgs {

			var hint TypeHint
			err := json.Unmarshal(d.Body, &hint)
			failOnError(err, "Failed to unmarshal type")

			notifier, factoryFunc, err := setup.GetNotifier(hint.Type)
			failOnError(err, "Failed to unmarshal payload")

			payload := factoryFunc()
			err = json.Unmarshal(d.Body, payload)

			failOnError(err, "Failed to unmarshal notification")

			err = notifier.Send(payload)
			failOnError(err, "Failed to send a message")
			d.Ack(false)
		}
	}()

	// Block forever (waiting for messages)
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

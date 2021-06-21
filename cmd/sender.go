package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func Fire(startDate int64, meterId string, mqAddress string) {
	data := genMeterData(startDate, meterId)
	meter_payload, err := json.Marshal(data)
	if err != nil {
		fmt.Errorf("Could not marshall struct into json")
	}
	// fmt.Print(string(meter_payload))
	sendMeterPayLoad(mqAddress, string(meter_payload))
}

// Sends payload to rabbitmq
func sendMeterPayLoad(mqAddress string, message string) {
	conn, err := amqp.Dial(mqAddress)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"meter-queue", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	failOnError(err, "Failed to declare a queue")
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(message),
		})
	log.Printf(" [x] Sent %s", message)
	failOnError(err, "Failed to publish a message")
}

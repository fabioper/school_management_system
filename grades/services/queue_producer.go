package services

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func PublishMessage(queueName string, data interface{}) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	handleError(err)
	defer conn.Close()

	ch, err := conn.Channel()
	handleError(err)
	defer ch.Close()

	_, err = ch.QueueDeclare(queueName, false, false, false, false, nil)
	handleError(err)

	body, err := json.Marshal(data)
	handleError(err)

	err = ch.Publish("", queueName, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        body,
	})
	fmt.Println("New message published!")
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}

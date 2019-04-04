package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func failOnError(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
	}
}
func main() {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	defer connection.Close()
	failOnError("Couldn't establish connection: ", err)
	channel, err := connection.Channel()
	_, err = channel.QueueDeclare("my_queue", true, false, false, false, nil)
	failOnError("Couldn't create queue", err)
	err = channel.QueueBind("my_queue", "#", "events", false, nil)
	failOnError("Couldn't bind queue", err)
	msgs, err := channel.Consume("my_queue", "", false, false, false, false, nil)
	failOnError("Error consuming queue", err)
	for msg := range msgs {
		fmt.Println("Message received: " + string(msg.Body))
		msg.Ack(false)
	}
}

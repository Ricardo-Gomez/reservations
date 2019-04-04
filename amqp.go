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
	failOnError("Error opening the channel", err)
	err = channel.ExchangeDeclare("events", "topic", false, false, false, false, nil)
	failOnError("Error declaring exchange", err)
	message := amqp.Publishing{
		Body: []byte("hello world"),
	}
	err = channel.Publish("events", "routekey", false, false, message)
	failOnError("msg string", err)
	fmt.Print("done")
}

package main

import (
	"github.com/htaidirt/go-rabbitmq-demo/lib/broker"
	"github.com/streadway/amqp"
)

func main() {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	consumer, err := broker.NewConsumer(connection)
	if err != nil {
		panic(err)
	}
	consumer.Listen()
}

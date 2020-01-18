package main

import (
	"fmt"
	"os"

	"github.com/htaidirt/go-rabbitmq-blog/lib/broker"
	"github.com/streadway/amqp"
)

func main() {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	producer, err := broker.NewProducer(connection)
	if err != nil {
		panic(err)
	}

	name := os.Args[1]
	producer.Push(fmt.Sprintf("Hello %s!", name))
}

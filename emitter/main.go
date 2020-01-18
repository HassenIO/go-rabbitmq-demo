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

	producer.Push(fmt.Sprintf("Pushing to broker with name=%s", os.Args[1]), os.Args[1])
}

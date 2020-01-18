package broker

import (
	"github.com/streadway/amqp"
)

var exchangeName = "GREETINGS"

func useQueue(ch *amqp.Channel) (amqp.Queue, error) {
	// QueueDeclare declares a queue to hold messages and deliver to consumers.
	// Declaring creates a queue if it doesn't already exist, or ensures that an existing queue matches the same parameters.
	// Source: https://godoc.org/github.com/streadway/amqp#Channel.QueueDeclare
	return ch.QueueDeclare(
		"",    // Nameless queue. In this case, the server will generate a unique name which will be returned in the Name field of Queue struct.
		false, // Not durable
		false, // Non auto-delete
		true,  // Exclusive queue (only accessible by the connection that declares them and will be deleted when the connection closes.)
		false, // No noWait. When noWait is true, the queue will assume to be declared on the server. A channel exception will arrive if the conditions are met for existing queues or attempting to modify an existing queue from a different connection.
		nil,   // No other args
	)
	// Non-Durable and Non-Auto-Deleted queues will remain declared as long as the server is running regardless of how many consumers.
	// This lifetime is useful for temporary topologies that may have long delays between consumer activity.
	// These queues can only be bound to non-durable exchanges.
}

func useExchange(ch *amqp.Channel) error {
	// ExchangeDeclare declares an exchange on the server.
	// If the exchange does not already exist, the server will create it.
	// If the exchange exists, the server verifies that it is of the provided type, durability and auto-delete flags.
	// Source: https://godoc.org/github.com/streadway/amqp#Channel.ExchangeDeclare
	return ch.ExchangeDeclare(
		exchangeName, // The exchange name
		"topic",      // Defines how messages are routed through the exchange ("direct", "fanout", "topic" and "headers")
		true,         // Durable messages
		false,        // Do not auto-delete
		false,        // Do accept publishing
		false,        // Do wait confirmation from the server
		nil,          // No other args
	)
	// Durable and Non-Auto-Deleted exchanges will survive server restarts and remain declared when there are no remaining bindings.
	// This is the best lifetime for long-lived exchange configurations like stable routes and default exchanges.
}

package broker

import (
	"log"

	"github.com/htaidirt/go-rabbitmq-blog/lib/logger"
	"github.com/streadway/amqp"
)

// Consumer struct type
type Consumer struct {
	connection *amqp.Connection
	queueName  string
}

func (c *Consumer) init() error {
	ch, err := c.connection.Channel()
	logger.OnError(err, "Error creating connectiion to channel")
	return useExchange(ch)
}

// Listen listens to the channel publications and print them to the console
func (c *Consumer) Listen(topics []string) error {
	ch, err := c.connection.Channel()
	logger.OnError(err, "Error listening to topics")
	defer ch.Close()

	q, err := useQueue(ch)
	logger.OnError(err, "Error creating queue")

	// For each topic, create a queue binding
	for _, t := range topics {
		// QueueBind binds an exchange to a queue so that publishings to the exchange will be routed
		// to the queue when the publishing routing key matches the binding routing key.
		err = ch.QueueBind(
			q.Name,       // Queue name
			t,            // Topic
			exchangeName, // Exchange namee
			false,        // Do wait
			nil,          // No other args
		)
		logger.OnError(err, "Error binding queue on new topic")
	}

	msg, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	logger.OnError(err, "Error consuming a new message")

	foreverLoop := make(chan bool) // Will never be filled, blocking the server from terminating
	go func() {
		for m := range msg {
			log.Printf("Received a message: %s", m.Body)
		}
	}()
	log.Printf("-> Listening for messages on Exchange=%s, and Queue=%s\n", exchangeName, q.Name)
	log.Printf("-> ctrl-c to exit")
	<-foreverLoop

	return nil
}

// NewConsumer create a new consumer with valid initalization
func NewConsumer(connection *amqp.Connection) (Consumer, error) {
	consumer := Consumer{
		connection: connection,
	}

	err := consumer.init()
	if err != nil {
		return Consumer{}, err
	}

	return consumer, nil
}

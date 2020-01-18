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

func (c *Consumer) init() error {
	ch, err := c.connection.Channel()
	logger.OnError(err, "Error creating a connection of a channel for consumer")
	defer ch.Close()

	return useExchange(ch)
}

// Listen listens to the channel publications and print them to the console
func (c *Consumer) Listen() error {
	ch, err := c.connection.Channel()
	logger.OnError(err, "Error creating connection to channel")
	defer ch.Close()

	q, err := useQueue(ch)
	logger.OnError(err, "Error creating queue")

	// QueueBind binds an exchange to a queue so that publishings to the exchange will be routed
	// to the queue when the publishing routing key matches the binding routing key.
	err = ch.QueueBind(
		q.Name,       // Queue name
		"",           // All topics
		exchangeName, // Exchange name
		false,        // Do wait (noWait=false)
		nil,          // No other args
	)
	logger.OnError(err, "Error binding queue in Listen")

	// Begin receiving on the returned chan Delivery before any other operation on the Connection or Channel.
	msg, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	logger.OnError(err, "Error consuming a new message")

	stop := make(chan bool) // Will be used to block the termination of the Listen function (must use ctrl-c)
	go func() {
		for m := range msg {
			log.Printf("Received a message: %s", m.Body)
		}
	}()
	log.Printf("-> Listening for messages on Exchange=%s and Queue=%s\n", exchangeName, q.Name)
	log.Printf("-> ctrl-c to exit")
	<-stop

	return nil
}

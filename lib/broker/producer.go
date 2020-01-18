package broker

import (
	"github.com/htaidirt/go-rabbitmq-blog/lib/logger"
	"github.com/streadway/amqp"
	"log"
)

// Producer type
type Producer struct {
	connection *amqp.Connection
}

func (p *Producer) init() error {
	channel, err := p.connection.Channel()
	logger.OnError(err, "Error setting up connection for channel")
	defer channel.Close()

	return useExchange(channel)
}

// Push the message in the exchange
func (p *Producer) Push(message string, severity string) error {
	channel, err := p.connection.Channel()
	logger.OnError(err, "Error setting up connection for channel")
	defer channel.Close()

	err = channel.Publish(
		exchangeName,
		severity,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	log.Printf("Sending message: %s -> %s", message, exchangeName)
	return nil
}

// NewProducer creates a new publisher and ensures the connection is established to our AMQP server
func NewProducer(connection *amqp.Connection) (Producer, error) {
	publisher := Producer{
		connection: connection,
	}

	err := publisher.init()
	if err != nil {
		return Producer{}, err
	}

	return publisher, nil
}

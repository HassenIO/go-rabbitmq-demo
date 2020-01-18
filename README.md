# Go ♥ RabbitMQ

A demo of communicating microservices written in Go with RabbitMQ.

Read the full story on [https://htaidirt.com/communicating-go-microservices-with-rabbitmq](https://htaidirt.com/communicating-go-microservices-with-rabbitmq)

## The short story

We create 2 microservices written in Go that should communicate through RabbitMQ. One receives a name as a parameter of the CLI and publishes a message to the RabbitMQ queue. The second should greet the person with the message sent by the publisher.

This is intended to be an easy demo. You can change the emitter to receive a name through HTTP, and the receiver to write some data in a database, or send an email.

## Executing the code

3 services will be launched through the terminal:

- **Launch RabbitMQ** by running `make rabbitmq` in your terminal. This will run a Docker container with RabbitMQ and opens port 5672.
- **Listen to messages to show** by running `make listen` in your terminal. This will run the Go receiver microservice that listen to greeting messages from RabbitMQ.
- **Send a name to greet** by running `make greet name=xxx` in your terminal. This will run the Go emitter microservice that sends to greeting message with the name (xxx) to RabbitMQ.

# License

MIT of course.

(but any reference to this page or the blog post is welcome)

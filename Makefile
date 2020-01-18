rabbitmq:
	docker run --name rabbitmq -p 5672:5672 rabbitmq:3.8
listen:
	go run receiver/main.go log.WARN
greeting:
	go run emitter/main.go $(name)
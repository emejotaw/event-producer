package main

import (
	"github.com/emejotaw/event-producer/config"
	"github.com/emejotaw/event-producer/internal/routes"
	"github.com/emejotaw/event-producer/pkg/events/rabbitmq"
)

func main() {

	cfg, err := config.LoadConfig()

	if err != nil {
		panic(err)
	}

	rabbitCfg := &cfg.Event.Producer.RabbitMQ
	rabbitMQ, err := rabbitmq.NewRabbitMQ(rabbitCfg.Username, rabbitCfg.Password, rabbitCfg.Host, rabbitCfg.Port)

	if err != nil {
		panic(err)
	}

	router := routes.NewRouter()
	router.Start(":3000", rabbitMQ)
}

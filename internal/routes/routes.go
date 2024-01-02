package routes

import (
	"github.com/emejotaw/event-producer/internal/controller"
	"github.com/emejotaw/event-producer/pkg/events"
	"github.com/gofiber/fiber/v2"
)

type router struct {
}

func NewRouter() *router {

	return &router{}
}

func (r *router) Start(port string, eventHandler events.EventHandler) {

	app := fiber.New()
	eventController := controller.NewEventController(eventHandler)
	app.Post("/api/v1/events", eventController.Produce)
	app.Listen(port)
}

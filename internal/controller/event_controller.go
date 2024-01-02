package controller

import (
	"net/http"

	"github.com/emejotaw/event-producer/internal/service"
	"github.com/emejotaw/event-producer/pkg/dto"
	"github.com/emejotaw/event-producer/pkg/events"
	"github.com/gofiber/fiber/v2"
)

type EventController struct {
	eventService *service.EventService
}

func NewEventController(eventHandler events.EventHandler) *EventController {

	eventService := service.NewEventService(eventHandler)

	return &EventController{
		eventService: eventService,
	}
}

func (ec *EventController) Produce(c *fiber.Ctx) error {

	eventDTO := &dto.EventDTO{}
	err := c.BodyParser(eventDTO)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return err
	}

	err = ec.eventService.Produce(eventDTO)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return err
	}

	c.Status(http.StatusAccepted)
	return nil
}

package service

import (
	"encoding/json"
	"log"

	"github.com/emejotaw/event-producer/pkg/dto"
	"github.com/emejotaw/event-producer/pkg/events"
)

type EventService struct {
	eventHandler events.EventHandler
}

func NewEventService(eventHandler events.EventHandler) *EventService {
	return &EventService{
		eventHandler: eventHandler,
	}
}

func (es *EventService) Produce(eventDTO *dto.EventDTO) error {

	dataBytes, err := json.Marshal(eventDTO)

	if err != nil {
		log.Printf("could not parse body, error: %v", err)
		return err
	}

	err = es.eventHandler.Publish(dataBytes)

	if err != nil {
		log.Printf("could not publish the event, error: %v", err)
		return err
	}

	log.Printf("Event published sucessfully")

	return nil
}

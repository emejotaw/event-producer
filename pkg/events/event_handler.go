package events

type EventHandler interface {
	Publish(body []byte) error
}

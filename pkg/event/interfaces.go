package event

type EventBusInterface interface {
	Publish(event Event)
	Subscribe() <-chan Event
}

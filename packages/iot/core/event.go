package core

type Event[T any] struct {
	Type    EventType
	Payload T
}

type EventType string

const (
	PowerOn EventType = "power_on"
)

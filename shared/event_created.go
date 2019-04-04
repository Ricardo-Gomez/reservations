package shared

import "time"

// se dispara cada que un evento se reserva
type EventCreatedEvent struct {
	ID         string
	Name       string
	LocationID string
	Start      time.Time
	End        time.Time
}

//nombre del evento que el message broker encola
func (e *EventCreatedEvent) EventName() string {
	return "eventCreated"
}

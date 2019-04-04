package shared

// se dispara cada que un evento se reserva
type EventBookedEvent struct {
	EventID string
	UserID  string
}

//nombre del evento que el message broker encola
func (e *EventBookedEvent) EventName() string {
	return "eventBooked"
}

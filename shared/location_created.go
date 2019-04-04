package shared

// se dispara cada que un evento se reserva
type LocationCreatedEvent struct {
	ID      string
	Name    string
	Address string
	Country string
	Hall    []string
}

//nombre del evento que el message broker encola
func (e *LocationCreatedEvent) EventName() string {
	return "locationCreated"
}

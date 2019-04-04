package msgqueue

//definicion para los eventos que son disparados usando los eventemiiters
type Event interface {
	EventName() string
}

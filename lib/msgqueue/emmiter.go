package msgqueue

//definicion de lo que tienen que implementar los eventemmiters
type EventEmmiter interface {
	Emmit(e Event) error
}

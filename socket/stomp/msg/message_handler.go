package msg

type Handler interface {
	HandleMessage(msg Message[interface{}]) error
}

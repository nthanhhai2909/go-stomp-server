package socket

import (
	"mem-ws/socket/core/stomp/stompmsg"
)

type SendingOperations[P interface{}] interface {
	Send(destination string, message stompmsg.Message[P]) error
}

package stomp

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"mem-ws/core/stomp/msg"
	"time"
)

const (
	writeWait = 10 * time.Second
)

type BrokerMessageHandler[T []byte] struct {
	Conn     *websocket.Conn
	Outbound chan T
	UserID   string
}

func NewBrokerMessageHandler(conn *websocket.Conn) msg.Handler[[]byte] {
	broker := &BrokerMessageHandler[[]byte]{
		Conn:     conn,
		Outbound: make(chan []byte),
		UserID:   uuid.New().String(),
	}
	return broker
}

func (broker BrokerMessageHandler[T]) HandleMessage(msg msg.Message[T]) error {
	return nil
}

func (broker BrokerMessageHandler[T]) GetUserID() string {
	return broker.UserID
}

func (broker BrokerMessageHandler[T]) GetConn() *websocket.Conn {
	return broker.Conn
}

func (broker BrokerMessageHandler[T]) GetOutboundChannel() chan T {
	return broker.Outbound
}

func (broker *BrokerMessageHandler[T]) outbound() {
	fmt.Println("Client Outbound InboundChannel started listen")

	defer func() {
		err := broker.Conn.Close()
		if err != nil {
			return
		}
	}()

	for {
		select {
		case payload, ok := <-broker.Outbound:
			err := broker.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err != nil {
				// TODO HGA WILL ADAPT LATER
			}

			if !ok {
				err = broker.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					// TODO HGA WILL ADAPT LATER
				}
				return
			}

			w, err := broker.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			_, err = w.Write(payload)
			if err != nil {
				// TODO HGA WILL ADAPT LATER
			}

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}

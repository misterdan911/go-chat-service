package hub

import (
	"fmt"
	"github.com/gofiber/contrib/websocket"
)

type Message struct {
	Name  string
	Price string
}

// Hub maintains the set of active clients and broadcasts messages to the clients
type Hub struct {
	Clients               map[*websocket.Conn]bool
	ClientRegisterChannel chan *websocket.Conn
	ClientRemovalChannel  chan *websocket.Conn
	BroadcastMessage      chan Message
}

// NewHub initializes a new Hub
func NewHub() *Hub {
	return &Hub{
		Clients:               make(map[*websocket.Conn]bool),
		ClientRegisterChannel: make(chan *websocket.Conn),
		ClientRemovalChannel:  make(chan *websocket.Conn),
		BroadcastMessage:      make(chan Message),
	}
}

// Run listens on channels and processes client registration, removal, and message broadcasting
func (h *Hub) Run() {
	fmt.Println("Hub is running")
	for {
		select {
		case conn := <-h.ClientRegisterChannel:
			h.Clients[conn] = true
		case conn := <-h.ClientRemovalChannel:
			delete(h.Clients, conn)
		case msg := <-h.BroadcastMessage:
			for conn := range h.Clients {
				_ = conn.WriteJSON(msg)
			}
		}
	}
}

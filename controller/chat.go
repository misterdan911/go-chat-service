package controller

import (
	"github.com/gofiber/contrib/websocket"
	"go-chat-service/internal/hub"
)

func HandleChat(h *hub.Hub) func(*websocket.Conn) {
	//fmt.Println("masuk BidPrice()")
	return func(conn *websocket.Conn) {
		//fmt.Println("masuk BidPrice anon func()")
		defer func() {
			h.ClientRemovalChannel <- conn

		}()

		name := conn.Query("name", "anonymous")
		h.ClientRegisterChannel <- conn

		for {
			messageType, price, err := conn.ReadMessage()
			//fmt.Println("masuk BidPrice anon func() ReadMessage")
			if err != nil {
				return
			}

			if messageType == websocket.TextMessage {
				//fmt.Println("masuk BidPrice anon func() ReadMessage > TextMessage")
				h.BroadcastMessage <- hub.Message{
					Name:  name,
					Price: string(price),
				}
			}
		}

	}
}

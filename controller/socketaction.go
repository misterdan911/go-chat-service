package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/contrib/socketio"
)

// MessageObject Basic chat message object
type MessageObject struct {
	Data  string `json:"data"`
	From  string `json:"from"`
	Event string `json:"event"`
	To    string `json:"to"`
}

func DefineSocketAction(clients *map[string]string) {
	// Multiple event handling supported
	socketio.On(socketio.EventConnect, func(ep *socketio.EventPayload) {
		fmt.Printf("Connection event 1 - User: %s", ep.Kws.GetStringAttribute("user_id"))
	})

	// Custom event handling supported
	socketio.On("CUSTOM_EVENT", func(ep *socketio.EventPayload) {
		fmt.Printf("Custom event - User: %s", ep.Kws.GetStringAttribute("user_id"))
		// --->

		// DO YOUR BUSINESS HERE

		// --->
	})
	// Custom event handling supported
	socketio.On("MAKAN_MAKAN", func(ep *socketio.EventPayload) {
		fmt.Printf("Custom event MAKAN_MAKAN - User: %s", ep.Kws.GetStringAttribute("user_id"))
		// --->

		// DO YOUR BUSINESS HERE

		// --->
	})

	// On message event
	socketio.On(socketio.EventMessage, func(ep *socketio.EventPayload) {

		fmt.Printf("Message event - User: %s - Message: %s", ep.Kws.GetStringAttribute("user_id"), string(ep.Data))

		message := MessageObject{}

		// Unmarshal the json message
		// {
		//  "from": "<user-id>",
		//  "to": "<recipient-user-id>",
		//  "event": "CUSTOM_EVENT",
		//  "data": "hello"
		//}
		err := json.Unmarshal(ep.Data, &message)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Fire custom event based on some
		// business logic
		if message.Event != "" {
			ep.Kws.Fire(message.Event, []byte(message.Data))
		}

		// Emit the message directly to specified user
		err = ep.Kws.EmitTo((*clients)[message.To], ep.Data, socketio.TextMessage)
		if err != nil {
			fmt.Println(err)
		}
	})

	// On disconnect event
	socketio.On(socketio.EventDisconnect, func(ep *socketio.EventPayload) {
		// Remove the user from the local clients
		delete(*clients, ep.Kws.GetStringAttribute("user_id"))
		fmt.Printf("Disconnection event - User: %s", ep.Kws.GetStringAttribute("user_id"))
	})

	// On close event
	// This event is called when the server disconnects the user actively with .Close() method
	socketio.On(socketio.EventClose, func(ep *socketio.EventPayload) {
		// Remove the user from the local clients
		delete(*clients, ep.Kws.GetStringAttribute("user_id"))
		fmt.Printf("Close event - User: %s", ep.Kws.GetStringAttribute("user_id"))
	})

	// On error event
	socketio.On(socketio.EventError, func(ep *socketio.EventPayload) {
		fmt.Printf("Error event - User: %s", ep.Kws.GetStringAttribute("user_id"))
	})

}

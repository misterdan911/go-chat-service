package mysocketio

import (
	"fmt"
	"github.com/gofiber/contrib/socketio"
)

// The key for the map is message.to
var Clients map[string]string
var Kws *socketio.Websocket

func DefineSocketAction() {

	// Multiple event handling supported
	socketio.On(socketio.EventConnect, func(ep *socketio.EventPayload) {
		fmt.Println("Connection event - User: ", ep.Kws.GetStringAttribute("user_id"))
	})

}

package mysocketio

import (
	"fmt"
	"github.com/gofiber/contrib/socketio"
)

// The key for the map is message.to
var Clients map[string]string

func DefineSocketAction() {

	// Multiple event handling supported
	socketio.On(socketio.EventConnect, func(ep *socketio.EventPayload) {
		fmt.Printf("Connection event 1 - User: %s", ep.Kws.GetStringAttribute("user_id"))
	})

}

package endpoints

import (
	"log"

	"github.com/gofiber/contrib/websocket"
)

var (
	Players    = make(map[*websocket.Conn]struct{})
	Register   = make(chan *websocket.Conn)
	Unregister = make(chan *websocket.Conn)
	Broadcast  = make(chan string)
)

func GetSocket(c *websocket.Conn) {

	defer func() {
		Unregister <- c
		c.Close()
	}()

	Register <- c

	for {
		messageType, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			return
		}

		// Maybe do something

		if messageType == websocket.TextMessage {
			Broadcast <- string(message)
		} else {
			log.Println("message received of type: ", messageType)
		}
	}

}

// Recieving command from the connection
func SocketListener() {
	for {
		select {

		case connection := <-Register:
			Players[connection] = struct{}{}
			log.Println("connection is registerd")

		case message := <-Broadcast:
			log.Println("message recieved:", message)

			// Iterate through each player and send back their own message
			for connection := range Players {
				err := connection.WriteMessage(websocket.TextMessage, []byte(message))

				if err != nil {
					log.Println("write error:", err)

					Unregister <- connection
					connection.WriteMessage(websocket.CloseMessage, []byte{})
					connection.Close()
				}
			}

		case connection := <-Unregister:
			delete(Players, connection)
			log.Println("connection is unregistered")
		}
	}

}

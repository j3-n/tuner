package endpoints

import (
	"encoding/json"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/j3-n/tuner/api/internal/models"
)

var (
	Players    = make(map[*websocket.Conn]struct{})
	Register   = make(chan *websocket.Conn)
	Unregister = make(chan *websocket.Conn)
	Broadcast  = make(chan models.Question)
)

func GetSocket(c *websocket.Conn) {

	defer func() {
		Unregister <- c
		c.Close()
	}()

	i := 0
	Register <- c

	for {

		messageType, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			return
		}
		if string(message) == "get" {
			if messageType == websocket.TextMessage {
				questions := models.QuestionsSet[i]
				i = i + 1
				Broadcast <- questions
			} else {
				log.Println("message received of type: ", messageType)
			}
		}

	}

}

// Recieving command from the connection
func SocketListener() {
	for {
		select {

		case connection := <-Register:
			Players[connection] = struct{}{}
			log.Println("connection is registered")

		case questions := <-Broadcast:
			log.Println("message received:", questions)

			// Encode the QuestionsSet to JSON
			jsonData, err := json.Marshal(questions)
			if err != nil {
				log.Println("error encoding QuestionsSet:", err)
				continue
			}

			for connection := range Players {
				err := connection.WriteMessage(websocket.TextMessage, jsonData)
				if err != nil {
					log.Println("write error:", err)
					Unregister <- connection
				}
			}

		case connection := <-Unregister:
			delete(Players, connection)
			log.Println("connection is unregistered")
		}
	}
}

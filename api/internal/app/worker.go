package app

import (
	"fmt"

	"github.com/gofiber/contrib/websocket"
	"github.com/j3-n/tuner/api/internal/models"
)

func PlayerWorker(c *websocket.Conn, p *models.Player, l *models.Lobby) {
	// Continuously poll for messages from the client
	defer c.Close()
	for {
		_, m, err := c.ReadMessage()
		if err != nil {
			// Player disconnect
			fmt.Printf("%s has disconnected from lobby %s\n", p.DisplayName, l.LobbyId)
			break
		}
		// Handle message
		if string(m) == "hello" {
			c.WriteMessage(websocket.TextMessage, []byte("fuck off"))
		}
	}

	// Cleanup

}

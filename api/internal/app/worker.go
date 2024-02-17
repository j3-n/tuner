package app

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/contrib/websocket"
	"github.com/j3-n/tuner/api/internal/models"
)

func PlayerWorker(c *websocket.Conn, p *models.Player, l *models.Lobby) {
	// Continuously poll for messages from the client
	for {
		_, m, err := c.ReadMessage()
		if err != nil {
			// Player disconnect
			fmt.Printf("%s has disconnected from lobby %s\n", p.DisplayName, l.LobbyId)
			break
		}
		// Handle message
		if string(m) == "quiz" {
			data, _ := json.Marshal(l.GenerateQuiz(10))
			c.WriteMessage(websocket.TextMessage, data)
		}
	}

	// Cleanup
	l.RemovePlayer(p)
	if len(l.PlayerList) == 0 {
		lobbies.RemoveLobby(l)
	} else {
		// Rebroadcast lobby
		data, _ := json.Marshal(l)
		l.BroadcastToAllPlayers(data)
	}
}

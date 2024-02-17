package app

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gofiber/contrib/websocket"
	"github.com/j3-n/tuner/api/internal/game"
	"github.com/j3-n/tuner/api/internal/models"
)

type Shit struct {
	Optype string `json:"command"` // Contains type of operation - GAME START, GUESS ANSWER ETC
	Data   string `json:"body"`    // Contains data relating to option above
}

func PlayerWorker(c *websocket.Conn, p *models.Player, l *models.Lobby) {
	// Continuously poll for messages from the client
	for {
		var fuck Shit
		err := c.ReadJSON(&fuck)
		if err != nil {
			// Player disconnect
			fmt.Printf("%s has disconnected from lobby %s\n", p.DisplayName, l.LobbyId)
			break
		}

		if fuck.Optype == "START" {
			// START GAME AT PLAYERS LOBBY
			//lobbies.Get(lobby).STARTGAME()
			// So set state to start game
			// broadcast to all other
			game.StartRound(l)
		} else if fuck.Optype == "GUESS" {
			// Data will contain id of answer
			// Store guess data into channel with data type map[player.uuid]answerIndex. Channel should be stored in Lobby
			type DonkeyBalls struct {
				AnswerId string `json:"answerId`
			}
			var cock DonkeyBalls
			c.ReadJSON(cock)
			answerInt, err := strconv.Atoi(cock.AnswerId)
			if err != nil {
				fmt.Println("Error converting data from guess")
			}
			l.Guesses[p.UUID] = answerInt
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

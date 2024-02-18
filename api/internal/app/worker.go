package app

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/contrib/websocket"
	"github.com/j3-n/tuner/api/internal/game"
	"github.com/j3-n/tuner/api/internal/models"
)

type ClientData struct {
	Optype string `json:"command"` // Contains type of operation - GAME START, GUESS ANSWER ETC
	Data   any    `json:"body"`    // Contains data relating to option above
}

type GuessData struct {
	AnswerId string `json:"answerId"`
}

func PlayerWorker(c *websocket.Conn, p *models.Player, l *models.Lobby) {
	// Continuously poll for messages from the client
	for {
		var data ClientData
		err := c.ReadJSON(&data)
		fmt.Println(data)
		if err != nil {
			// Player disconnect
			fmt.Printf("%s has disconnected from lobby %s\n", p.DisplayName, l.LobbyId)
			break
		}

		if data.Optype == "START" {
			// START GAME AT PLAYERS LOBBY
			//lobbies.Get(lobby).STARTGAME()
			// So set state to start game
			// broadcast to all other
			game.StartRound(l)
		} else if data.Optype == "GUESS" {
			// Data will contain id of answer
			// Store guess data into channel with data type map[player.uuid]answerIndex. Channel should be stored in Lobby
			_, guessed := l.Guesses[p.UUID]
			if l.State == models.Guessing && !guessed {

				valMap, ok := data.Data.(map[string]any)
				if !ok {
					log.Println("Error occurred")
				}
				val, ok := valMap["answerId"]
				if !ok {
					log.Println("Error occurred")
				}
				valStr, ok := val.(string)
				if !ok {
					log.Println("Error occurred")
				}
				answerInt, err := strconv.Atoi(valStr)
				if err != nil {
					log.Println("Error converting data from guess,", err)
				}
				l.Guesses[p.UUID] = answerInt
				// TODO: check if all players have submitted a guess
			}
		}
	}

	// Cleanup
	l.RemovePlayer(p)
	if len(l.PlayerList) == 0 {
		lobbies.RemoveLobby(l)
	} else {
		// Rebroadcast lobby
		l.BroadcastToAllPlayers(models.CreatePacket(l.State.String(), l))
	}
}

package app

import (
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/j3-n/tuner/api/internal/game"
	"github.com/j3-n/tuner/api/internal/models"
)

func PlayerWorker(c *websocket.Conn, p *models.Player, l *models.Lobby) {
	// Continuously poll for messages from the client

	Timer := time.NewTimer(time.Second * 15)
	// Stop the timer at the end of the function.
	// Defers are called when the parent function exits.
	defer Timer.Stop()

	for {
		type Shit struct {
			Optype string // Contains type of operation - GAME START, GUESS ANSWER ETC
			Data   string // Contains data relating to option above
		}
		var fuck Shit
		c.ReadJSON(fuck)

		if fuck.Optype == "START" {
			// START GAME AT PLAYERS LOBBY
			//lobbies.Get(lobby).STARTGAME()
			// So set state to start game
			// broadcast to all other
			game.StartRound(l)
		} else if fuck.Optype == "GUESS" {

			// Data wilsl contain id of answer
			// Send this data to nathans function which will evaluate it when round is over
			// nathans function should wait till all people in lobby have given an answer
			now := time.Now().Add(15 * time.Second)
			<-Timer.C
			game.WaitingGuesses(l, now)
		}
	}

	// Cleanup
	l.RemovePlayer(p)
}

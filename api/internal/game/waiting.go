package game

import (
	"time"

	"github.com/j3-n/tuner/api/internal/models"
)

func WaitingGuesses(l *models.Lobby, times time.Time) {
	// Loop indefinitely
	for l.State == models.Guessing {
		now := time.Now()
		differnce := now.Sub(times)
		err := []byte(differnce.Abs().String())
		l.BroadcastToAllPlayers(err)

		if differnce == 0 {
			l.State = models.Results
			return

		}

		if len(l.Guesses) == len(l.PlayerList) {
			l.State = models.Results
			return
		}
	}
}

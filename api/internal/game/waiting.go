package game

import (
	"time"

	"github.com/j3-n/tuner/api/internal/models"
)

func WaitingGuesses(l *models.Lobby, times time.Time) {
	ticker := time.NewTicker(time.Second)
	// Loop indefinitely
	for {
		select {
		case <-ticker.C:
			now := time.Now()
			differnce := now.Sub(times)
			err := []byte(differnce.Abs().String())
			l.BroadcastToAllPlayers(err)
		}
	}
}

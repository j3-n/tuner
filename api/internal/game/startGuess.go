package game

import (
	"encoding/json"

	"github.com/j3-n/tuner/api/internal/models"
)

func StartRound(l *models.Lobby) {
	// Broadcast questions
	if l.State == models.Waiting {
		message, err := json.Marshal(models.QuestionsSet)
		if err != nil {
			return
		}
		l.BroadcastToAllPlayers(message)

		l.State = models.Guessing
	}
}

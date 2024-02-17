package game

import (
	"encoding/json"

	"github.com/j3-n/tuner/api/internal/models"
)

func StartRound(l *models.Lobby, playerID int, guess string) {
	// Broadcast that game has started
	shit := []byte("START")
	l.BroadcastToAllPlayers(shit)

	message, err := json.Marshal(models.QuestionsSet)
	if err != nil {
		return
	}
	l.BroadcastToAllPlayers(message)

	// Send questions to frontend with guesses
	// Wait
	// Points -> Leaderboards
}

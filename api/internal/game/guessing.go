package game

import (
	"github.com/j3-n/tuner/api/internal/models"
)

func StartRound(l *models.Lobby, playerID int, guess string) {
	// Broadcast that game has started
	shit := []byte("START")
	l.BroadcastToAllPlayers(shit)

	
	// Send questions to frontend with guesses
	// Wait
	// Points -> Leaderboards
}

package game

import "github.com/j3-n/tuner/api/internal/models"

func WaitingGuesses(l *models.Lobby) {
	// Waiting until packet of (optype) is sent
	for l.State == models.Guessing {
		// Wait until everyone has submitted their shit or time runs out
	}
}

package game

import "github.com/j3-n/tuner/api/internal/models"

func Results(l *models.Lobby) {
	// Updates player scores
	// Loops through list of player scores
	// Send scores to frontend
	// Wait
	// Points -> Guessing
	for _, player := range l.PlayerList {
		index, ok := l.Guesses[player.UUID]
		if ok {
			if l.CurrentQuestion[l.Round].Correct == index {

			}
		}

	}
	l.Round++
}

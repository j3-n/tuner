package game

import (
	"encoding/json"
	"time"

	"github.com/j3-n/tuner/api/internal/models"
)

func StartRound(l *models.Lobby) {
	// Generate quiz
	if l.State == models.Waiting {
		l.Questions = l.GenerateQuiz(10)
		l.State = models.Guessing
		// Broadcast question
		m, _ := json.Marshal(l.Questions[0])
		l.BroadcastToAllPlayers(m)
		l.Guesses = map[string]int{}
		// Start timer for question deadline
		l.Timer = time.AfterFunc(time.Second*15, func() {
			EndRound(l)
		})
	}
}

func EndRound(l *models.Lobby) {

}

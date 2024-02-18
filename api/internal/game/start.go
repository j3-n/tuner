package game

import (
	"time"

	"github.com/j3-n/tuner/api/internal/models"
)

func StartRound(l *models.Lobby) {
	// Generate quiz
	if l.State == models.Waiting {
		l.Questions = l.GenerateQuiz(3)
		l.Points = map[string]int{}
		l.State = models.Guessing
		NewRound(l)
	}
}

func NewRound(l *models.Lobby) {
	// Broadcast question
	l.BroadcastToAllPlayers(models.CreatePacket(l.State.String(), l.Questions[l.Round]))
	l.Guesses = map[string]int{}
	// Start timer for question deadline
	l.Timer = time.AfterFunc(time.Second*15, func() {
		EndRound(l)
	})
}

func EndRound(l *models.Lobby) {
	l.State = models.Results
	Results(l)
	// Next round after 5 seconds
	time.AfterFunc(time.Second*5, func() {
		NextRound(l)
	})
}

func NextRound(l *models.Lobby) {
	l.Round += 1
	if l.Round >= len(l.Questions) {
		l.State = models.Finish
		// Game over
		GameOver(l)
	} else {
		l.State = models.Guessing
		// Next round
		NewRound(l)
	}
}

func GameOver(l *models.Lobby) {
	l.BroadcastToAllPlayers(models.CreatePacket(l.State.String(), l.CreateLeaderboard()))
}

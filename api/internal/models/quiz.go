package models

import (
	"context"
	"math/rand"

	"github.com/zmb3/spotify/v2"
)

type Answer struct {
	Artist   string `json:"artist"`
	Song     string `json:"song"`
	AnswerID int    `json:"answerId"`
}

type Question struct {
	Question string   `json:"question"`
	Answers  []Answer `json:"answers"`
	Correct  int      `json:"correct"`
}

// ONLY CALL THIS ONCE SPOTIFY API SLOW
func (l *Lobby) GenerateQuiz(length int) []Question {
	questions, answers := l.GenerateQuizSongs(length)
	// Convert to questions
	ret := []Question{}
	for _, q := range questions {
		correct := rand.Intn(3)
		qanswers := []Answer{}
		for i, a := range answers[:4] {
			qanswers = append(qanswers, Answer{
				AnswerID: i,
				Song:     a.Name,
				Artist:   a.Artists[0].Name,
			})
		}
		qanswers[correct] = Answer{
			Artist:   q.Artists[0].Name,
			Song:     q.Name,
			AnswerID: correct,
		}
		question := Question{
			Question: q.PreviewURL,
			Answers:  qanswers,
			Correct:  correct,
		}
		ret = append(ret, question)

	}
	rand.Shuffle(len(answers), func(i int, j int) {
		temp := answers[i]
		answers[i] = answers[j]
		answers[j] = temp
	})
	return ret
}

func (l *Lobby) GenerateQuizSongs(length int) ([]*spotify.FullTrack, []*spotify.FullTrack) {
	// in honour of Stephen Piddock
	songsExist := map[string]bool{}
	allSongs := []*spotify.FullTrack{}
	l.mu.Lock()
	defer l.mu.Unlock()

	for _, p := range l.PlayerList {
		// Fetch top songs from each player and push to map
		songs, err := p.Client.CurrentUsersTopTracks(context.Background())
		if err != nil {
			continue
		}
		// Append songs
		for _, s := range songs.Tracks {
			if !songsExist[s.ID.String()] {
				songsExist[s.ID.String()] = true
				allSongs = append(allSongs, &s)
			}
		}
	}
	// Shuffle the songs
	length = max(length, len(allSongs))
	rand.Shuffle(len(allSongs), func(i int, j int) {
		temp := allSongs[i]
		allSongs[i] = allSongs[j]
		allSongs[j] = temp
	})
	return allSongs[:length], allSongs[length+1:]
}

var QuestionsSet = []Question{
	{
		Question: "Who is this artist",
		Answers: []Answer{
			{Artist: "Taylor Swift", Song: "hi", AnswerID: 1},
			{Artist: "Ye", Song: "good monring", AnswerID: 2},
			{Artist: "JayZ", Song: "heie", AnswerID: 3},
			{Artist: "SZA", Song: "Ddsds", AnswerID: 4},
		},
		Correct: 1,
	},
	{
		Question: "What is this song",
		Answers: []Answer{
			{Artist: "Taylor Swift", Song: "hi", AnswerID: 1},
			{Artist: "Ye", Song: "good monring", AnswerID: 2},
			{Artist: "JayZ", Song: "heie", AnswerID: 3},
			{Artist: "SZA", Song: "Ddsds", AnswerID: 4},
		},
		Correct: 2,
	},
}

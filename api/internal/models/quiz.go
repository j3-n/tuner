package models

type Answer struct {
	Artist   string `json:"artist"`
	Song     string `json:"song"`
	AnswerID int    `json:"answerId"`
}

type Questions struct {
	Question string   `json:"question"`
	Answers  []Answer `json:"answers"`
	Correct  int      `json:"correct"`
}

var QuestionsSet = []Questions{
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

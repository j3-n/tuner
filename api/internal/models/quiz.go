package models

type Answer struct {
	Artist   string `json:"artist"`
	Song     string `json:"song"`
	AnswerID int    `json:"id"`
}

type Questions struct {
	Question string   `json:"question"`
	Answers  []Answer `json:"answers"`
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
	},
}

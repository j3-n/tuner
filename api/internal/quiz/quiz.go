package quiz

type Answer struct {
	Answer   string `json:"answer"`
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
			{Answer: "Taylor Swift", AnswerID: 1},
			{Answer: "Ye", AnswerID: 2},
			{Answer: "JayZ", AnswerID: 3},
			{Answer: "SZA", AnswerID: 4},
		},
	},
}

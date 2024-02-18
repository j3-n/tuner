package models

type Competitor struct {
	Id    int `gorm:"primaryKey" json:"id"`
	Score int `gorm:"not null" json:"score"`
}

// type LeaderBoard struct {
// 	Comp []Competitor
// }

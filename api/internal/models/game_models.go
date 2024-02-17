package models

type Lobby struct {
	LobbyId    int `json:"lobbyId"`
	PlayerList []Player
}

type Player struct {
	PlayerId int    `json:"playerId"`
	Name     string `json:"playerName"`
}

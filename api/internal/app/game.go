package app

import "math/rand"

type Lobby struct {
	lobbyId    int `json:"lobbyId"`
	playerList []Player
}

type Player struct {
	playerId int    `json:"playerId"`
	name     string `json:"playerName"`
}

var gameLobbies []Lobby

// Creates lobby and returns lobby id
func CreateLobby() int {

	randomLobbyID := -1

	// Generate random id until unique one found ;) pls dont hurt me
	isUnique := false
	for isUnique == false {
		isUnique = true
		randomLobbyID = int(rand.Float64() * 1000)
		for _, lobby := range gameLobbies {
			if lobby.lobbyId == randomLobbyID {
				isUnique = false
			}
		}
	}

	gameLobbies = append(gameLobbies, Lobby{lobbyId: randomLobbyID})
	return randomLobbyID
}

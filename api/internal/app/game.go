package app

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/gofiber/contrib/websocket"
)

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
func CreateLobby(hostPlayer int) int {

	randomLobbyID := -1

	// Generate random id until unique one found ;) pls dont hurt me
	isUnique := false
	for !isUnique {
		isUnique = true
		randomLobbyID = int(rand.Float64() * 1000)
		for _, lobby := range gameLobbies {
			if lobby.lobbyId == randomLobbyID {
				isUnique = false
			}
		}
	}

	gameLobbies = append(gameLobbies, Lobby{lobbyId: randomLobbyID})
	AddPlayerToLobby(Player{playerId: hostPlayer, name: "John"}, randomLobbyID)
	return randomLobbyID
}

// Adds player to lobby with provided player id and lobby id
func AddPlayerToLobby(player Player, lobbyID int) error {
	for lobbyIndex, lobby := range gameLobbies {
		if lobby.lobbyId == lobbyID {
			flag := false
			// Check that player with given id doesnt exist in lobby
			for _, playerId := range lobby.playerList {
				if playerId.playerId == player.playerId {
					flag = true
				}
			}
			if !flag {
				gameLobbies[lobbyIndex].playerList = append(gameLobbies[lobbyIndex].playerList, player)
			} else {
				return errors.New("Cannot add duplicate player " + player.name)
			}
		}
	}
	return nil
}

func HandleCreationRequest(c *websocket.Conn) {
	receivedPlayerId, err := strconv.Atoi(c.Params("playerId"))
	if err != nil {
		fmt.Printf("Error parsing player id: %s\n", c.Params("playerId"))
		return
	}
	fmt.Printf("Received params: %d\n", receivedPlayerId)
	// Authenticate player first
	CreateLobby(receivedPlayerId)
}

package app

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"sync"

	"github.com/gofiber/contrib/websocket"
	"github.com/j3-n/tuner/api/internal/models"
	"golang.org/x/oauth2"
)

type User struct {
	uuid  string
	token *oauth2.Token
}

type Users struct {
	users []*User
	mu    *sync.Mutex
}

var gameLobbies []models.Lobby
var users Users

func (u *Users) Add(n *User) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.users = append(u.users, n)
}

// Creates lobby and returns lobby id using bogo lobby algorithm
func CreateLobby(hostPlayer int) int {

	randomLobbyID := -1

	// Generate random id until unique one found ;) pls dont hurt me
	isUnique := false
	for !isUnique {
		isUnique = true
		randomLobbyID = int(rand.Float64() * 1000)
		for _, lobby := range gameLobbies {
			if lobby.LobbyId == randomLobbyID {
				isUnique = false
			}
		}
	}

	gameLobbies = append(gameLobbies, models.Lobby{LobbyId: randomLobbyID})
	AddPlayerToLobby(models.Player{PlayerId: hostPlayer, Name: "John"}, randomLobbyID)
	return randomLobbyID
}

// Adds player to lobby with provided player id and lobby id
func AddPlayerToLobby(player models.Player, lobbyID int) error {
	for lobbyIndex, lobby := range gameLobbies {
		if lobby.LobbyId == lobbyID {
			flag := false
			// Check that player with given id doesnt exist in lobby
			for _, playerId := range lobby.PlayerList {
				if playerId.PlayerId == player.PlayerId {
					flag = true
				}
			}
			if !flag {
				gameLobbies[lobbyIndex].PlayerList = append(gameLobbies[lobbyIndex].PlayerList, player)
			} else {
				return errors.New("Cannot add duplicate player " + player.Name)
			}
		}
	}
	return nil
}

// Handle websocket request for lobby creation
func HandleCreationRequest(c *websocket.Conn) {
	receivedPlayerId, err := strconv.Atoi(c.Params("playerId"))
	if err != nil {
		fmt.Printf("Error parsing player id: %s\n", c.Params("playerId"))
		return
	}
	fmt.Printf("Received params: %d\n", receivedPlayerId)
	// Check player is actually connected

	// Create lobby
	CreateLobby(receivedPlayerId)
}

// Reads player id and lobby id and assigns player to lobby if both player id and lobby exist
func HandleAddPlayerRequest(c *websocket.Conn) {

}

func HandlePlayerGuess(c *websocket.Conn) {

	type shit struct {
		Uuid     string `json:"uuid"`
		AnswerId int    `json:"answerId"`
	}
	var thing shit

	c.ReadJSON(&thing)

	// Update score

	// Send back correct or nah
}

package app

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	"github.com/gofiber/contrib/websocket"
	"github.com/j3-n/tuner/api/internal/models"
	"github.com/zmb3/spotify/v2"
)

var gameLobbies []models.Lobby
var users models.Users

// Creates lobby and returns lobby id using bogo lobby algorithm
func CreateLobby() int {

	randomLobbyID := -1

	// Generate random id until unique one found ;) pls dont hurt me
	isUnique := false
	for !isUnique {
		isUnique = true
		randomLobbyID = int(rand.Float64() * 1000)
		for _, lobby := range gameLobbies {
			if lobby.LobbyId == fmt.Sprintf("%d", randomLobbyID) {
				isUnique = false
			}
		}
	}

	gameLobbies = append(gameLobbies, models.Lobby{LobbyId: fmt.Sprintf("%d", randomLobbyID)})
	return randomLobbyID
}

// Adds player to lobby with provided player id and lobby id
func AddPlayerToLobby(player models.Player, lobbyID string) error {
	for lobbyIndex, lobby := range gameLobbies {
		if lobby.LobbyId == lobbyID {
			flag := false
			// Check that player with given id doesnt exist in lobby
			for _, playerId := range lobby.PlayerList {
				if playerId.UUID == player.UUID {
					flag = true
				}
			}
			if !flag {
				gameLobbies[lobbyIndex].PlayerList = append(gameLobbies[lobbyIndex].PlayerList, player)
			} else {
				return errors.New("Cannot add duplicate player " + player.DisplayName)
			}
		}
	}
	return nil
}

// Handle websocket request for lobby creation
func HandleCreationRequest(c *websocket.Conn) {
	// Create lobby
	id := CreateLobby()
	if JoinLobby(c, fmt.Sprintf("%d", id)) == nil {
		c.WriteMessage(websocket.TextMessage, []byte("yippee lobby"))
	}
}

func HandleAddPlayerRequest(c *websocket.Conn) {
	id := c.Params("lobby")
	if JoinLobby(c, id) == nil {
		c.WriteMessage(websocket.TextMessage, []byte("yippee"))
	}
}

func JoinLobby(c *websocket.Conn, lobby string) error {
	// Check player authentication
	p := CreatePlayer(c.Cookies("TUNER_SESSION"))
	if p == nil {
		return errors.New("User not authenticated!")
	}
	// Join lobby
	fmt.Printf("%s has connected via websocket to lobby %s\n", p.DisplayName, lobby)
	return nil
}

func CreatePlayer(uuid string) *models.Player {
	// Check that the Spotify token still works
	if len(uuid) == 0 || !users.Exists(uuid) {
		return nil
	}
	client := spotify.New(auth.Client(context.Background(), users.Get(uuid).Token))
	u, err := client.CurrentUser(context.Background())
	if err != nil {
		return nil
	}
	return &models.Player{
		User:        users.Get(uuid),
		Client:      client,
		DisplayName: u.DisplayName,
	}
}

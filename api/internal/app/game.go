package app

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"

	"github.com/gofiber/contrib/websocket"
	"github.com/j3-n/tuner/api/internal/models"
	"github.com/zmb3/spotify/v2"
)

var lobbies models.Lobbies
var users models.Users

// Creates lobby and returns lobby id using bogo lobby algorithm
func CreateLobby() int {

	randomLobbyID := -1

	// Generate random id until unique one found ;) pls dont hurt me
	isUnique := false
	for !isUnique {
		randomLobbyID = int(rand.Float64() * 1000)
		isUnique = !lobbies.Exists(fmt.Sprintf("%d", randomLobbyID))
	}

	lobbies.Add(&models.Lobby{LobbyId: fmt.Sprintf("%d", randomLobbyID)})
	return randomLobbyID
}

// Adds player to lobby with provided player id and lobby id
func AddPlayerToLobby(player *models.Player, lobbyID string) error {
	lobby := lobbies.Get(lobbyID)
	if lobby == nil {
		return errors.New("invalid lobby")
	}
	for _, playerId := range lobby.PlayerList {
		if playerId.UUID == player.UUID {
			return errors.New("player is already in this lobby")
		}
	}
	lobby.PlayerList = append(lobby.PlayerList, player)
	return nil
}

// Handle websocket request for lobby creation
func HandleCreationRequest(c *websocket.Conn) {
	defer c.Close()
	// Create lobby
	id := CreateLobby()
	JoinLobby(c, fmt.Sprintf("%d", id))
}

// Reads player id and lobby id and assigns player to lobby if both player id and lobby exist
func HandleAddPlayerRequest(c *websocket.Conn) {
	defer c.Close()
	id := c.Params("lobby")
	JoinLobby(c, id)
}

func JoinLobby(c *websocket.Conn, lobby string) {
	// Check player authentication
	p := CreatePlayer(c.Cookies("TUNER_SESSION"))
	if p == nil {
		c.WriteMessage(websocket.TextMessage, []byte("Not authenticated!"))
		return
	}
	// Join lobby
	err := AddPlayerToLobby(p, lobby)
	if err != nil {
		c.WriteMessage(websocket.TextMessage, []byte("Invalid lobby!"))
		return
	}
	fmt.Printf("%s is joining lobby %s\n", p.DisplayName, lobby)
	// Send lobby information as JSON
	l, _ := json.Marshal(lobbies.Get(lobby))
	c.WriteMessage(websocket.TextMessage, []byte(l))
	// Send to running worker
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

// Add/Update question in lobby when new game is started and when new round is started
func AddQuestionToLobby(question models.Questions, lobby models.Lobby) {
	for _, lob := range gameLobbies {
		if lob.LobbyId == lobby.LobbyId {
			lobby.CurrentQuestion = question
		}
	}
}

func HandlePlayerGuess(c *websocket.Conn) {

	type shit struct {
		Uuid     string `json:"uuid"`
		AnswerId int    `json:"answerId"`
	}
	var thing shit

	c.ReadJSON(&thing)

	fmt.Println(thing)
	// Update score

	// Send back correct or nah
}

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
	if lobby.HasPlayer(player) {
		return errors.New("player already in lobby")
	}
	lobby.AddPlayer(player)
	return nil
}

// Handle websocket request for lobby creation
func HandleCreationRequest(c *websocket.Conn) {
	// Create lobby
	defer c.Close()
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
	p := CreatePlayer(c, c.Cookies("TUNER_SESSION"))
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
	// Send lobby information as JSON to all connected players
	lo := lobbies.Get(lobby)
	l, _ := json.Marshal(lo)
	lo.BroadcastToAllPlayers(l)
	// Send to running worker
	PlayerWorker(c, p, lo)
}

func CreatePlayer(c *websocket.Conn, uuid string) *models.Player {
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
		Conn:        c,
	}

}

// Listener for each player
func MonitorPlayer(c *websocket.Conn, p models.Player, lobby string) {
	type Shit struct {
		Optype string // Contains type of operation - GAME START, GUESS ANSWER ETC
		Data   string // Contains data relating to option above
	}
	var fuck Shit
	c.ReadJSON(fuck)

	if fuck.Optype == "START" {
		// START GAME AT PLAYERS LOBBY
		//lobbies.Get(lobby).STARTGAME()
		// So set state to start game
		// broadcast to all other
	} else if fuck.Optype == "GUESS" {
		// Data will contain id of answer
		// Send this data to nathans function which will evaluate it when round is over
		// nathans function should wait till all people in lobby have given an answer
	}

}

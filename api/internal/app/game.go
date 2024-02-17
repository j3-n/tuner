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
	lob := models.Lobby{LobbyId: fmt.Sprintf("%d", randomLobbyID), State: models.Waiting, Round: 0}
	lob.CurrentQuestion = lob.GenerateQuiz(4)
	lobbies.Add(&lob)
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
	go PlayerWorker(c, p, lo)
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
		IconURL:     u.Images[0].URL,
		Carousel:    GenerateCarousel(client),
		Conn:        c,
	}

}

func GenerateCarousel(c *spotify.Client) []string {
	// Fetch user's top albums to use in carousel
	songs, err := c.CurrentUsersTopTracks(context.Background())
	if err != nil {
		return nil
	}
	exists := map[string]bool{}
	carousel := []string{}
	for _, s := range songs.Tracks {
		if !exists[s.Album.ID.String()] {
			exists[s.Album.ID.String()] = true
			carousel = append(carousel, s.Album.Images[0].URL)
		}
	}
	return carousel
}

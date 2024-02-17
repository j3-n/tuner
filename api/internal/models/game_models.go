package models

import (
	"sync"

	"github.com/gofiber/contrib/websocket"
	"github.com/zmb3/spotify/v2"
	"golang.org/x/oauth2"
)

type Lobby struct {
	LobbyId         string    `json:"lobbyId"`
	PlayerList      []*Player `json:"players"`
	CurrentQuestion Questions
}

type User struct {
	UUID  string
	Token *oauth2.Token
}

type Users struct {
	users []*User
	mu    sync.Mutex
}

type Lobbies struct {
	lobbies []*Lobby
	mu      sync.Mutex
}

type Player struct {
	*User       `json:"-"`
	Client      *spotify.Client `json:"-"`
	DisplayName string          `json:"displayName"`
	Conn        websocket.Conn
}

func (l *Lobbies) Add(lo *Lobby) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.lobbies = append(l.lobbies, lo)
}

func (l *Lobbies) Exists(id string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	for _, lo := range l.lobbies {
		if lo.LobbyId == id {
			return true
		}
	}

	return false
}

func (l *Lobbies) Get(id string) *Lobby {
	l.mu.Lock()
	defer l.mu.Unlock()
	for _, lo := range l.lobbies {
		if lo.LobbyId == id {
			return lo
		}
	}
	return nil
}

func (u *Users) Add(n *User) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.users = append(u.users, n)
}

func (u *Users) Exists(uuid string) bool {
	u.mu.Lock()
	defer u.mu.Unlock()

	for _, us := range u.users {
		if us.UUID == uuid {
			return true
		}
	}
	return false
}

func (u *Users) Get(uuid string) *User {
	u.mu.Lock()
	defer u.mu.Unlock()
	for _, us := range u.users {
		if us.UUID == uuid {
			return us
		}
	}

	return nil
}

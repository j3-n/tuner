package models

import (
	"sync"

	"github.com/zmb3/spotify/v2"
	"golang.org/x/oauth2"
)

type Lobby struct {
	LobbyId    string `json:"lobbyId"`
	PlayerList []Player
}

type User struct {
	UUID  string
	Token *oauth2.Token
}

type Users struct {
	users []*User
	mu    sync.Mutex
}

type Player struct {
	*User
	Client      *spotify.Client
	DisplayName string
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
	for _, us := range u.users {
		if us.UUID == uuid {
			return us
		}
	}

	return nil
}

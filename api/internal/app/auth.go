package app

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/zmb3/spotify/v2"
)

func Auth(c *fiber.Ctx) error {
	r, err := adaptor.ConvertRequest(c, false)
	if err != nil {
		return err
	}

	token, err := auth.Token(c.Context(), state, r)
	if err != nil {
		return err
	}

	// Create spotify client from token
	client := spotify.New(auth.Client(c.Context(), token))

	page, _ := client.CurrentUsersTopTracks(c.Context())

	for i, track := range page.Tracks {
		fmt.Printf("%d: %s - %s (%s)\n", i+1, track.Name, track.Artists[0].Name, track.Album.Name)
	}

	return nil
}

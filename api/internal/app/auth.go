package app

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/zmb3/spotify/v2"
)

func Auth(c *fiber.Ctx) error {
	ctx := context.Background()

	r, err := adaptor.ConvertRequest(c, false)
	if err != nil {
		return err
	}

	token, err := auth.Token(ctx, state, r)
	if err != nil {
		return err
	}

	// Create spotify client from token
	client := spotify.New(auth.Client(ctx, token))

	page, _ := client.CurrentUsersTopTracks(ctx)

	for i, track := range page.Tracks {
		fmt.Printf("%d: %s - %s (%s)\n", i+1, track.Name, track.Artists[0].Name, track.PreviewURL)
	}

	return nil
}

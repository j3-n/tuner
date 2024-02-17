package app

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/google/uuid"
	"github.com/j3-n/tuner/api/internal/models"
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

	// Create token for the client
	s := uuid.New()
	// Send the token as a cookie
	cookie := new(fiber.Cookie)
	cookie.Name = "TUNER_SESSION"
	cookie.Value = s.String()
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.Cookie(cookie)

	// Register the player
	users.Add(&models.User{
		UUID:  s.String(),
		Token: token,
	})

	fmt.Printf("%s registered\n", s.String())

	return nil
}

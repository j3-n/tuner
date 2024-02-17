package app

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
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

	fmt.Println(token)
	return nil
}

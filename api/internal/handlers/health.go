package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetHealth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	}
}

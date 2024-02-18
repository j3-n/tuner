package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/j3-n/tuner/api/internal/database"
	"github.com/j3-n/tuner/api/internal/models"
)

func GetPlayers(s *database.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var players [10]models.Competitor

		err := s.List(&players, "competitors")
		if err != nil {
			// Handles error when dishes cannot be retrieved
			return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving Players" + err.Error())
		}
		// Sends menu as JSON if no errors have occured
		return c.JSON(players)
	}
}

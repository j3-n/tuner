package handlers

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/j3-n/tuner/api/internal/database"
	"github.com/j3-n/tuner/api/internal/models"
)

func ListPlayers(s *database.Store) fiber.Handler {
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

func GetPlayer(s *database.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		pk, err := c.ParamsInt("id")
		if err != nil {
			// Handles error when players cannot be retrieved
			return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving Players")
		}
		var players models.Competitor

		// Gets players from table
		err = s.Get(&players, "competitors", pk)
		if err != nil {
			// Handles error when players cannot be retrieved
			return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving Players")
		}

		// Sends menu as JSON if no errors have occured
		return c.JSON(players)
	}
}

func PostPlayer(s *database.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var players models.Competitor
		err := c.BodyParser(&players)
		log.Println(err)
		if err != nil {
			// Handles error when dplayers cannot be retrieved
			return c.Status(fiber.StatusInternalServerError).SendString("Error adding player")
		}

		err = s.Add(&players, "competitors")
		if err != nil {
			// Handles error when players cannot be retrieved
			return c.Status(fiber.StatusInternalServerError).SendString("Error adding player")
		}

		return c.SendStatus(http.StatusNoContent)
	}
}

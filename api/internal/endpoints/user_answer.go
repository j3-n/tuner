package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"github.com/j3-n/tuner/api/internal/models"
)

// Gets answer from user
func PostUserAnswer(c *fiber.Ctx) error {
	var bodyData models.Answer
	if err := c.BodyParser(&bodyData); err != nil {
		return err
	}
	return c.JSON(bodyData)
}

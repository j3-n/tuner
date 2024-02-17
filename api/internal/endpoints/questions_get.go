package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"github.com/j3-n/tuner/api/internal/models"
)

func GetQuestions(c *fiber.Ctx) error {
	eer := models.QuestionsSet[0]
	return c.JSON(eer)
}

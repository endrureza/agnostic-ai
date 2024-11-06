package handlers

import (
	"github.com/endrureza/agnostic-ai/internal/services"
	"github.com/endrureza/agnostic-ai/pkg/models"
	"github.com/gofiber/fiber/v2"
)

type AIHandler struct {
	aiService *services.AIService
}

func NewAIHandler(aiService *services.AIService) *AIHandler {
	return &AIHandler{
		aiService: aiService,
	}
}

func (h *AIHandler) GenerateText(c *fiber.Ctx) error {
	req := new(models.GenerateTextRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	text, err := h.aiService.GenerateText(*req)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(text)
}

func (h *AIHandler) GenerateImage(c *fiber.Ctx) error {
	req := new(models.GenerateImageRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	image, err := h.aiService.GenerateImage(*req)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(image)
}

func (h *AIHandler) TranscribedAudio(c *fiber.Ctx) error {
	req := new(models.TranscribedAudioRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	text, err := h.aiService.TranscribedAudio(*req)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(text)
}

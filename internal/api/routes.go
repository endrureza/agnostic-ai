package api

import (
	"github.com/endrureza/agnostic-ai/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, aiHandler *handlers.AIHandler) {
	app.Post("/chat", aiHandler.GenerateText)
	app.Post("/image", aiHandler.GenerateImage)
	app.Post("/audio", aiHandler.TranscribedAudio)
}

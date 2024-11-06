package main

import (
	"github.com/endrureza/agnostic-ai/config"
	"github.com/endrureza/agnostic-ai/internal/api"
	"github.com/endrureza/agnostic-ai/internal/handlers"
	"github.com/endrureza/agnostic-ai/internal/providers"
	"github.com/endrureza/agnostic-ai/internal/services"
	"github.com/endrureza/agnostic-ai/pkg/models"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Load .env file
	cfg, err := config.LoadConfig()

	if err != nil {
		// Handle the error, e.g., log it and exit
		panic(err)
	}

	providerCfg := models.ProvidersConfig{
		OpenAI: config.OpenAIConfig{
			URL: cfg.OpenAI.URL,
			Key: cfg.OpenAI.Key,
		},
		Gemini: config.GeminiConfig{
			URL: cfg.Gemini.URL,
			Key: cfg.Gemini.Key,
		},
		Claude: config.ClaudeConfig{
			URL: cfg.Claude.URL,
			Key: cfg.Claude.Key,
		},
	}

	providers.RegisterProviders(providerCfg)

	aiService := services.NewAIService()

	aiHandler := handlers.NewAIHandler(aiService)

	api.RegisterRoutes(app, aiHandler)

	app.Listen("127.0.0.1:" + cfg.App.Port)
}

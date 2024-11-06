package models

import (
	"github.com/endrureza/agnostic-ai/config"
)

type ProvidersConfig struct {
	OpenAI config.OpenAIConfig `json:"openai"`
	Gemini config.GeminiConfig `json:"gemini"`
	Claude config.ClaudeConfig `json:"claude"`
}

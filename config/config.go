package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Env  string `json:"env"`
	Port string `json:"port"`
}

type OpenAIConfig struct {
	URL string `json:"url"`
	Key string `json:"key"`
}

type GeminiConfig struct {
	URL string `json:"url"`
	Key string `json:"key"`
}

type ClaudeConfig struct {
	URL string `json:"url"`
	Key string `json:"key"`
}

type Config struct {
	App    AppConfig    `json:"app"`
	OpenAI OpenAIConfig `json:"openai"`
	Gemini GeminiConfig `json:"gemini"`
	Claude ClaudeConfig `json:"claude"`
}

func LoadConfig() (Config, error) {
	// Load .env file
	err := godotenv.Load()

	if err != nil {
		return Config{}, fmt.Errorf("error loading .env file: %w", err)
	}

	cnf := Config{
		App: AppConfig{
			Env:  os.Getenv("APP_ENV"),
			Port: os.Getenv("APP_PORT"),
		},
		OpenAI: OpenAIConfig{
			URL: os.Getenv("OPENAI_BASE_URL"),
			Key: os.Getenv("OPENAI_API_KEY"),
		},
		Gemini: GeminiConfig{
			URL: os.Getenv("GEMINI_BASE_URL"),
			Key: os.Getenv("GEMINI_API_KEY"),
		},
		Claude: ClaudeConfig{
			URL: os.Getenv("CLAUDE_BASE_URL"),
			Key: os.Getenv("CLAUDE_API_KEY"),
		},
	}

	return cnf, nil
}

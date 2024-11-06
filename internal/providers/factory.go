package providers

import (
	"errors"

	"github.com/endrureza/agnostic-ai/pkg/models"
)

var providers = map[string]Provider{}

func RegisterProviders(cfg models.ProvidersConfig) {
	providers["openai"] = NewOpenAIProvider(cfg.OpenAI.Key, cfg.OpenAI.URL)
	providers["gemini"] = NewGeminiProvider(cfg.Gemini.Key, cfg.Gemini.URL)
	providers["claude"] = NewClaudeProvider(cfg.Claude.Key, cfg.Claude.URL)
}

func GetProvider(name string) (Provider, error) {
	provider, exists := providers[name]

	if !exists {
		return nil, errors.New("Provider not found")
	}

	return provider, nil
}

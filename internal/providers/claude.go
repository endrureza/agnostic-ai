package providers

import (
	"encoding/json"
	"fmt"

	"github.com/endrureza/agnostic-ai/pkg/models"
	"github.com/gofiber/fiber/v2"
)

type ClaudeProvider struct {
	Key string
	URL string
}

func NewClaudeProvider(Key string, URL string) *ClaudeProvider {
	return &ClaudeProvider{
		Key: Key,
		URL: URL,
	}
}

func (p *ClaudeProvider) GenerateText(req models.GenerateTextRequest) (models.GenerateTextResponse, error) {
	maxTokens := 2048

	if req.MaxTokens != 0 {
		maxTokens = req.MaxTokens
	}

	newBody := models.ClaudeTextRequest{
		Model:     req.Model,
		Messages:  req.Messages,
		MaxTokens: maxTokens,
	}

	requestBody, err := json.Marshal(newBody)

	if err != nil {
		return models.GenerateTextResponse{}, err
	}

	agent := fiber.Post(p.URL + "/messages")
	agent.Set("Content-Type", "application/json")
	agent.Set("x-api-key", p.Key)
	agent.Set("anthropic-version", "2023-06-01")
	agent.Body(requestBody)

	statusCode, body, errs := agent.Bytes()

	if len(errs) > 0 {
		return models.GenerateTextResponse{}, errs[0]
	}

	if statusCode != fiber.StatusOK {
		return models.GenerateTextResponse{}, fmt.Errorf("status code: %d", statusCode)
	}

	res := new(models.ClaudeTextResponse)

	if err := json.Unmarshal(body, res); err != nil {
		return models.GenerateTextResponse{}, err
	}

	newRes := models.GenerateTextResponse{
		Text: res.GetContent(),
	}

	return newRes, nil
}

func (p *ClaudeProvider) GenerateImage(req models.GenerateImageRequest) (models.GenerateImageResponse, error) {
	return models.GenerateImageResponse{}, fmt.Errorf("unimplemented")
}

func (p *ClaudeProvider) TranscribedAudio(req models.TranscribedAudioRequest) (models.TranscribedAudioResponse, error) {
	return models.TranscribedAudioResponse{}, fmt.Errorf("unimplemented")
}

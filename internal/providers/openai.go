package providers

import (
	"encoding/json"
	"fmt"

	"github.com/endrureza/agnostic-ai/pkg/models"
	"github.com/gofiber/fiber/v2"
)

type OpenAIProvider struct {
	Key string
	URL string
}

func NewOpenAIProvider(Key string, URL string) *OpenAIProvider {
	return &OpenAIProvider{
		Key: Key,
		URL: URL,
	}
}

func (p *OpenAIProvider) GenerateText(req models.GenerateTextRequest) (models.GenerateTextResponse, error) {
	if req.Temperature == 0 {
		req.Temperature = 0.3
	}

	if req.MaxTokens == 0 {
		req.MaxTokens = 2560
	}

	newBody := models.OpenAITextRequest{
		Model:       req.Model,
		Messages:    req.Messages,
		Temperature: req.Temperature,
		MaxTokens:   req.MaxTokens,
		Stream:      req.Stream,
	}

	requestBody, err := json.Marshal(newBody)

	if err != nil {
		return models.GenerateTextResponse{}, err
	}

	agent := fiber.Post(p.URL + "/chat/completions")
	agent.Set("Content-Type", "application/json")
	agent.Set("Authorization", "Bearer "+p.Key)
	agent.Body(requestBody)
	statusCode, body, errs := agent.Bytes()

	if len(errs) > 0 {
		return models.GenerateTextResponse{}, errs[0]
	}

	if statusCode != fiber.StatusOK {
		return models.GenerateTextResponse{}, fmt.Errorf("error: %d", statusCode)
	}

	res := new(models.OpenAITextResponse)

	if err := json.Unmarshal(body, &res); err != nil {
		return models.GenerateTextResponse{}, err
	}

	newRes := models.GenerateTextResponse{
		Text: res.GetContent(),
	}

	return newRes, nil
}

func (p *OpenAIProvider) GenerateImage(req models.GenerateImageRequest) (models.GenerateImageResponse, error) {
	if req.Size == "" {
		req.Size = "1024x1024"
	}

	if req.N == 0 {
		req.N = 1
	}

	newBody := models.OpenAIImageRequest{
		Model:  req.Model,
		Size:   req.Size,
		N:      req.N,
		Prompt: req.Prompt,
	}

	requestBody, err := json.Marshal(newBody)

	if err != nil {
		return models.GenerateImageResponse{}, err
	}

	agent := fiber.Post(p.URL + "/images/generations")
	agent.Set("Content-Type", "application/json")
	agent.Set("Authorization", "Bearer "+p.Key)
	agent.Body(requestBody)

	statusCode, body, errs := agent.Bytes()

	var decodeBody map[string]interface{}

	if err := json.Unmarshal(body, &decodeBody); err != nil {
		return models.GenerateImageResponse{}, err
	}

	fmt.Println(decodeBody)

	if len(errs) > 0 {
		return models.GenerateImageResponse{}, errs[0]
	}

	if statusCode != fiber.StatusOK {
		return models.GenerateImageResponse{}, fmt.Errorf("error: %d", statusCode)
	}

	res := new(models.OpenAIImageResponse)

	if err := json.Unmarshal(body, &res); err != nil {
		return models.GenerateImageResponse{}, err
	}

	newRes := models.GenerateImageResponse{
		Image: res.GetUrl(),
	}

	return newRes, nil
}

func (p *OpenAIProvider) TranscribedAudio(req models.TranscribedAudioRequest) (models.TranscribedAudioResponse, error) {
	requestBody, err := json.Marshal(req)

	if err != nil {
		return models.TranscribedAudioResponse{}, err
	}

	agent := fiber.Post(p.URL + "/audio/transcriptions")
	agent.Set("Content-Type", "multipart/form-data")
	agent.Set("Authorization", "Bearer "+p.Key)
	agent.Body(requestBody)

	statusCode, body, errs := agent.Bytes()

	if len(errs) > 0 {
		return models.TranscribedAudioResponse{}, errs[0]
	}

	if statusCode != fiber.StatusOK {
		return models.TranscribedAudioResponse{}, fmt.Errorf("error: %d", statusCode)
	}

	res := new(models.TranscribedAudioResponse)

	if err := json.Unmarshal(body, &res); err != nil {
		return models.TranscribedAudioResponse{}, err
	}

	return *res, nil
}
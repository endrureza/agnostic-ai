package providers

import (
	"encoding/json"
	"fmt"

	"github.com/endrureza/agnostic-ai/pkg/models"
	"github.com/gofiber/fiber/v2"
)

type HFPovider struct {
	Key string
	URL string
}

func NewHFProvider(Key string, URL string) *HFPovider {
	return &HFPovider{
		Key: Key,
		URL: URL,
	}
}

func (p *HFPovider) GenerateText(req models.GenerateTextRequest) (models.GenerateTextResponse, error) {
	newBody := models.HFTextRequest{
		Model:     req.Model,
		Messages:  req.Messages,
		MaxTokens: req.MaxTokens,
		Stream:    req.Stream,
	}

	requestBody, err := json.Marshal(newBody)

	if err != nil {
		return models.GenerateTextResponse{}, err
	}

	postUrl := p.URL + "/models/" + req.Model + "/v1/chat/completions"

	agent := fiber.Post(postUrl)
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

	res := new(models.HFTextResponse)

	if err := json.Unmarshal(body, &res); err != nil {
		return models.GenerateTextResponse{}, err
	}

	newRes := models.GenerateTextResponse{
		Text: res.GetContent(),
	}

	return newRes, nil
}

func (p *HFPovider) GenerateImage(req models.GenerateImageRequest) (models.GenerateImageResponse, error) {
	newBody := models.HFImageRequest{
		Inputs: req.Prompt,
	}

	requestBody, err := json.Marshal(newBody)

	if err != nil {
		return models.GenerateImageResponse{}, err
	}

	postUrl := p.URL + "/models/" + req.Model

	agent := fiber.Post(postUrl)
	agent.Set("Content-Type", "application/json")
	agent.Set("Authorization", "Bearer "+p.Key)
	agent.Body(requestBody)

	statusCode, body, errs := agent.Bytes()

	if len(errs) > 0 {
		return models.GenerateImageResponse{}, errs[0]
	}

	if statusCode != fiber.StatusOK {
		return models.GenerateImageResponse{}, fmt.Errorf("error: %d", statusCode)
	}

	res := new(models.HFImageResponse)

	if err := json.Unmarshal(body, &res); err != nil {
		return models.GenerateImageResponse{}, err
	}

	newRes := models.GenerateImageResponse{
		Image: res.Image,
	}

	return newRes, nil
}

func (p *HFPovider) TranscribedAudio(req models.TranscribedAudioRequest) (models.TranscribedAudioResponse, error) {
	return models.TranscribedAudioResponse{}, fmt.Errorf("unimplemented")
}

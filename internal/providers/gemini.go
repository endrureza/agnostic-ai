package providers

import (
	"encoding/json"
	"fmt"

	"github.com/endrureza/agnostic-ai/pkg/models"
	"github.com/gofiber/fiber/v2"
)

type GeminiProvider struct {
	Key string
	URL string
}

func NewGeminiProvider(Key string, URL string) *GeminiProvider {
	return &GeminiProvider{
		Key: Key,
		URL: URL,
	}
}

func (p *GeminiProvider) GenerateText(req models.GenerateTextRequest) (models.GenerateTextResponse, error) {
	newBody := models.GeminiTextRequest{
		Contents: []struct {
			Role  string `json:"role"`
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		}{},
	}

	for _, message := range req.Messages {
		newBody.Contents = append(newBody.Contents, struct {
			Role  string `json:"role"`
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		}{
			Role: message.Role,
			Parts: []struct {
				Text string `json:"text"`
			}{
				{Text: message.Content},
			},
		})
	}

	requestBody, err := json.Marshal(newBody)

	if err != nil {
		return models.GenerateTextResponse{}, err
	}

	postUrl := p.URL + "/models/" + req.Model + ":generateContent?key=" + p.Key

	if req.Stream {
		postUrl = p.URL + "/models/" + req.Model + ":streamGenerateContent?alt=sse&key=" + p.Key
	}

	agent := fiber.Post(postUrl)
	agent.Set("Content-Type", "application/json")
	agent.Body(requestBody)

	statusCode, body, errs := agent.Bytes()

	if len(errs) > 0 {
		return models.GenerateTextResponse{}, errs[0]
	}

	if statusCode != fiber.StatusOK {
		return models.GenerateTextResponse{}, fmt.Errorf("error: %d", statusCode)
	}

	res := new(models.GeminiTextResponse)

	if err := json.Unmarshal(body, &res); err != nil {
		return models.GenerateTextResponse{}, err
	}

	newRes := models.GenerateTextResponse{
		Text: res.GetContent(),
	}

	return newRes, nil
}
func (p *GeminiProvider) GenerateImage(req models.GenerateImageRequest) (models.GenerateImageResponse, error) {
	return models.GenerateImageResponse{}, fmt.Errorf("unimplemented")
}

func (p *GeminiProvider) TranscribedAudio(req models.TranscribedAudioRequest) (models.TranscribedAudioResponse, error) {
	return models.TranscribedAudioResponse{}, fmt.Errorf("unimplemented")
}

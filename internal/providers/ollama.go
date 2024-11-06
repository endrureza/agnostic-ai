package providers

import (
	"encoding/json"
	"fmt"

	"github.com/endrureza/agnostic-ai/pkg/models"
	"github.com/gofiber/fiber/v2"
)

type OllamaProvider struct {
	URL string
}

func NewOllamaProvider(URL string) *OllamaProvider {
	return &OllamaProvider{
		URL: URL,
	}
}

func (p *OllamaProvider) GenerateText(req models.GenerateTextRequest) (models.GenerateTextResponse, error) {
	newBody := models.OllamaTextRequest{
		Model:    req.Model,
		Messages: req.Messages,
		Stream:   req.Stream,
	}

	fmt.Println(p.URL + "/api/chat")

	requestBody, err := json.Marshal(newBody)

	if err != nil {
		return models.GenerateTextResponse{}, err
	}

	agent := fiber.Post(p.URL + "/api/chat")
	agent.Set("Content-Type", "application/json")
	agent.Body(requestBody)

	statusCode, body, errs := agent.Bytes()

	if len(errs) > 0 {
		return models.GenerateTextResponse{}, errs[0]
	}

	if statusCode != fiber.StatusOK {
		return models.GenerateTextResponse{}, fmt.Errorf("ollama responded with status code %d", statusCode)
	}

	res := new(models.OllamaTextResponse)

	if err := json.Unmarshal(body, &res); err != nil {
		return models.GenerateTextResponse{}, err
	}

	newRes := models.GenerateTextResponse{
		Text: res.Message.Content,
	}

	return newRes, nil
}

// GenerateImage implements Provider.
func (p *OllamaProvider) GenerateImage(req models.GenerateImageRequest) (models.GenerateImageResponse, error) {
	panic("unimplemented")
}

// TranscribedAudio implements Provider.
func (p *OllamaProvider) TranscribedAudio(req models.TranscribedAudioRequest) (models.TranscribedAudioResponse, error) {
	panic("unimplemented")
}

package providers

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/endrureza/agnostic-ai/pkg/models"
	"github.com/gofiber/fiber/v2"
)

type GroqPovider struct {
	Key string
	URL string
}

func NewGroqProvider(Key string, URL string) *GroqPovider {
	return &GroqPovider{
		Key: Key,
		URL: URL,
	}
}

func (p *GroqPovider) GenerateText(req models.GenerateTextRequest) (models.GenerateTextResponse, error) {
	newBody := models.GroqTextRequest{
		Model:       req.Model,
		Messages:    req.Messages,
		MaxTokens:   req.MaxTokens,
		Stream:      req.Stream,
		Temperature: req.Temperature,
	}

	requestBody, err := json.Marshal(newBody)

	if err != nil {
		return models.GenerateTextResponse{}, err
	}

	postUrl := p.URL + "/chat/completions"

	agent := fiber.Post(postUrl)
	agent.Set("Authorization", "Bearer "+p.Key)
	agent.Set("Content-Type", "application/json")
	agent.Body(requestBody)

	statusCode, body, errs := agent.Bytes()

	if len(errs) > 0 {
		return models.GenerateTextResponse{}, errs[0]
	}

	if statusCode != fiber.StatusOK {
		return models.GenerateTextResponse{}, fmt.Errorf("error: %d", statusCode)
	}

	res := new(models.GroqTextResponse)

	if err := json.Unmarshal(body, res); err != nil {
		return models.GenerateTextResponse{}, err
	}

	newRes := models.GenerateTextResponse{
		Text: res.GetContent(),
	}

	return newRes, nil
}

func (p *GroqPovider) GenerateImage(req models.GenerateImageRequest) (models.GenerateImageResponse, error) {
	return models.GenerateImageResponse{}, fmt.Errorf("unimplemented")
}

func (p *GroqPovider) TranscribedAudio(req models.TranscribedAudioRequest) (models.TranscribedAudioResponse, error) {
	src, err := req.File.Open()

	if err != nil {
		return models.TranscribedAudioResponse{}, err
	}

	defer src.Close()

	fileContent, err := io.ReadAll(src)

	if err != nil {
		return models.TranscribedAudioResponse{}, err
	}

	newBody := models.GroqAudioRequest{
		Model: req.Model,
		File:  fileContent,
	}

	args := fiber.AcquireArgs()
	args.Set("model", newBody.Model)

	agent := fiber.Post(p.URL + "/audio/transcriptions")
	agent.Set("Content-Type", "multipart/form-data")
	agent.Set("Authorization", "Bearer "+p.Key)
	agent.FileData(
		&fiber.FormFile{
			Name:      req.File.Filename,
			Fieldname: "file",
			Content:   newBody.File,
		},
	)
	agent.MultipartForm(args)

	fiber.ReleaseArgs(args)

	statusCode, body, errs := agent.Bytes()

	if len(errs) > 0 {
		return models.TranscribedAudioResponse{}, errs[0]
	}

	if statusCode != fiber.StatusOK {
		return models.TranscribedAudioResponse{}, fmt.Errorf("error: %d", statusCode)
	}

	res := new(models.GroqAudioResponse)

	if err := json.Unmarshal(body, &res); err != nil {
		return models.TranscribedAudioResponse{}, err
	}

	newRes := models.TranscribedAudioResponse{
		Text: res.Text,
	}

	return newRes, nil

	return models.TranscribedAudioResponse{}, fmt.Errorf("unimplemented")
}

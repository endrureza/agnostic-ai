package providers

import (
	"fmt"

	"github.com/endrureza/agnostic-ai/pkg/models"
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
	return models.GenerateTextResponse{}, fmt.Errorf("unimplemented")
}

func (p *GeminiProvider) GenerateImage(req models.GenerateImageRequest) (models.GenerateImageResponse, error) {
	return models.GenerateImageResponse{}, fmt.Errorf("unimplemented")
}

func (p *GeminiProvider) TranscribedAudio(req models.TranscribedAudioRequest) (models.TranscribedAudioResponse, error) {
	return models.TranscribedAudioResponse{}, fmt.Errorf("unimplemented")
}

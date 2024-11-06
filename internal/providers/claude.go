package providers

import (
	"fmt"

	"github.com/endrureza/agnostic-ai/pkg/models"
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
	return models.GenerateTextResponse{}, fmt.Errorf("unimplemented")
}

func (p *ClaudeProvider) GenerateImage(req models.GenerateImageRequest) (models.GenerateImageResponse, error) {
	return models.GenerateImageResponse{}, fmt.Errorf("unimplemented")
}

func (p *ClaudeProvider) TranscribedAudio(req models.TranscribedAudioRequest) (models.TranscribedAudioResponse, error) {
	return models.TranscribedAudioResponse{}, fmt.Errorf("unimplemented")
}

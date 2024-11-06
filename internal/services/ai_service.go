package services

import (
	"github.com/endrureza/agnostic-ai/internal/providers"
	"github.com/endrureza/agnostic-ai/pkg/models"
)

type AIService struct{}

func NewAIService() *AIService {
	return &AIService{}
}

func (s *AIService) GenerateText(req models.GenerateTextRequest) (models.GenerateTextResponse, error) {
	provider, err := providers.GetProvider(req.Provider)

	if err != nil {
		return models.GenerateTextResponse{}, err
	}

	return provider.GenerateText(req)
}

func (s *AIService) GenerateImage(req models.GenerateImageRequest) (models.GenerateImageResponse, error) {
	provider, err := providers.GetProvider(req.Provider)

	if err != nil {
		return models.GenerateImageResponse{}, err
	}

	return provider.GenerateImage(req)
}

func (s *AIService) TranscribedAudio(req models.TranscribedAudioRequest) (models.TranscribedAudioResponse, error) {
	provider, err := providers.GetProvider(req.Provider)

	if err != nil {
		return models.TranscribedAudioResponse{}, err
	}

	return provider.TranscribedAudio(req)
}

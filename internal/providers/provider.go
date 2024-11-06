package providers

import "github.com/endrureza/agnostic-ai/pkg/models"

type Provider interface {
	GenerateText(req models.GenerateTextRequest) (models.GenerateTextResponse, error)
	GenerateImage(req models.GenerateImageRequest) (models.GenerateImageResponse, error)
	TranscribedAudio(req models.TranscribedAudioRequest) (models.TranscribedAudioResponse, error)
}

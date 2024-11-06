package models

type GenerateTextResponse struct {
	Text string `json:"text"`
}

type GenerateImageResponse struct {
	Image string `json:"image"`
}

type TranscribedAudioResponse struct {
	Text string `json:"text"`
}

package models

import "mime/multipart"

type GenerateTextRequest struct {
	Provider string `json:"provider"`
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
	Temperature float64 `json:"temperature,omitempty"`
	MaxTokens   int     `json:"max_tokens,omitempty"`
	Stream      bool    `json:"stream,omitempty"`
}

type GenerateImageRequest struct {
	Provider string `json:"provider"`
	Model    string `json:"model"`
	Prompt   string `json:"prompt"`
	Size     string `json:"size,omitempty"`
	N        int    `json:"n,omitempty"`
}

type TranscribedAudioRequest struct {
	Provider string                `json:"provider"`
	Model    string                `json:"model"`
	File     *multipart.FileHeader `json:"file,omitempty"`
}

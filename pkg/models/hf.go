package models

type HFTextRequest struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
	MaxTokens int  `json:"max_tokens,omitempty" default:"2048"`
	Stream    bool `json:"stream,omitempty" default:"false"`
}

type HFTextResponse struct {
	ID                string `json:"id"`
	Created           int    `json:"created"`
	Model             string `json:"model"`
	SystemFingerprint string `json:"system_fingerprint"`
	Usage             struct {
		PromptTokens     int `json:"prompt_tokens"`
		TotalTokens      int `json:"total_tokens"`
		CompletionTokens int `json:"completion_tokens"`
	} `json:"usage"`
	Choices []struct {
		FinishReason string `json:"finish_reason"`
		Index        int    `json:"index"`
		Message      struct {
			Content string `json:"content"`
			Role    string `json:"role"`
			Refusal bool   `json:"refusal"`
		} `json:"message"`
		LogProbs float64 `json:"logprobs"`
	} `json:"choices"`
}

type HFImageRequest struct {
	Inputs     string `json:"inputs"`
	Parameters *struct {
		GuidanceScale     float64  `json:"guidance_scale,omitempty"`
		NegativePrompt    []string `json:"negative_prompt,omitempty"`
		NumInferenceSteps int      `json:"num_inference_steps,omitempty"`
		TargetSize        *struct {
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"target_size,omitempty"`
		Scheduler string `json:"scheduler,omitempty"`
		Seed      int    `json:"seed,omitempty"`
	} `json:"parameters,omitempty"`
}

type HFImageResponse struct {
	Image string `json:"image"`
}

type HFAudioRequest struct {
	Inputs     string                 `json:"inputs"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

type HFAudioResponse struct {
	Text   string `json:"text"`
	Chunks []struct {
		Text       string `json:"text"`
		Timestamps []int  `json:"timestamps"`
	} `json:"chunks"`
}

func (h HFTextResponse) GetContent() string {
	if len(h.Choices) > 0 && h.Choices[0].Message.Content != "" {
		return h.Choices[0].Message.Content
	}
	return ""
}

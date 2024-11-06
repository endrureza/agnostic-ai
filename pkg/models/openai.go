package models

type OpenAITextRequest struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
	MaxTokens   int     `json:"max_completion_tokens,omitempty"`
	Temperature float64 `json:"temperature,omitempty"`
	Stream      bool    `json:"stream,omitempty"`
}

type OpenAITextResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
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
	Usage struct {
		PromptTokens       int `json:"prompt_tokens"`
		CompletionTokens   int `json:"completion_tokens"`
		TotalTokens        int `json:"total_tokens"`
		PromptTokenDetails struct {
			CachedTokens int `json:"cached_tokens"`
		}
		CompletionTokenDetails struct {
			ReasoningTokens          int `json:"reasoning_tokens"`
			AcceptedPredictionTokens int `json:"accepted_prediction_tokens"`
			RejectedPredictionTokens int `json:"rejected_prediction_tokens"`
		}
	} `json:"usage"`
	SystemFingerprint string `json:"system_fingerprint"`
}

type OpenAIImageRequest struct {
	Prompt         string `json:"prompt"`
	Model          string `json:"model,omitempty"`
	N              int    `json:"n,omitempty"`
	Size           string `json:"size,omitempty"`
	Quality        int    `json:"quality,omitempty"`
	Style          string `json:"style,omitempty"`
	ResponseFormat string `json:"response_format,omitempty"`
}

type OpenAIImageResponse struct {
	Created int `json:"created"`
	Data    []struct {
		Url string `json:"url"`
	} `json:"data"`
}

type OpenAIAudioRequest struct {
	File                   []byte   `json:"file"`
	Model                  string   `json:"model"`
	Language               string   `json:"language,omitempty"`
	Prompt                 string   `json:"prompt,omitempty"`
	ResponseFormat         string   `json:"response_format,omitempty"`
	Temperature            float64  `json:"temperature,omitempty"`
	TimestampGranularities []string `json:"timestamp_granularities,omitempty"`
}

type OpenAIAudioResponse struct {
	Text string `json:"text"`
}

func (o OpenAITextResponse) GetContent() string {
	if len(o.Choices) > 0 && o.Choices[0].Message.Content != "" {
		return o.Choices[0].Message.Content
	}
	return ""
}

func (o OpenAIImageResponse) GetUrl() string {
	if len(o.Data) > 0 {
		return o.Data[0].Url
	}
	return ""
}

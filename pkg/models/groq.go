package models

type GroqTextRequest struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
	MaxTokens         int     `json:"max_tokens,omitempty"`
	PresencePenalty   float64 `json:"presence_penalty,omitempty"`
	FrequencyPenalty  float64 `json:"frequency_penalty,omitempty"`
	N                 int     `json:"n,omitempty"`
	ParallelToolCalls bool    `json:"parallel_tool_calls,omitempty"`
	ResponseFormat    *struct {
		Type string `json:"type"`
	} `json:"response_format,omitempty"`
	Seed          int      `json:"seed,omitempty"`
	Stop          []string `json:"stop,omitempty"`
	Stream        bool     `json:"stream,omitempty"`
	StreamOptions *struct {
		IncludeUsage bool `json:"include_usage"`
	} `json:"stream_options,omitempty"`
	Temperature float64 `json:"temperature,omitempty"`
	TopP        float64 `json:"top_p,omitempty"`
}

type GroqTextResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index        int    `json:"index"`
		Logprobs     bool   `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
		Message      struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Usage struct {
		QueueTime        float64 `json:"queue_time"`
		PromptTokens     int     `json:"prompt_tokens"`
		PrompTime        float64 `json:"prompt_time"`
		CompletionTokens int     `json:"completion_tokens"`
		CompletionTime   float64 `json:"completion_time"`
		TotalTokens      int     `json:"total_tokens"`
		TotalTime        float64 `json:"total_time"`
	} `json:"usage"`
	SystemFingerprint string `json:"system_fingerprint"`
	XGroq             struct {
		ID string `json:"id"`
	} `json:"x_groq"`
}

type GroqAudioRequest struct {
	File                   []byte   `json:"file"`
	Language               string   `json:"language,omitempty"`
	Model                  string   `json:"model"`
	Prompt                 string   `json:"prompt,omitempty"`
	ResponseFormat         string   `json:"response_format,omitempty"`
	Temperature            float64  `json:"temperature,omitempty"`
	TimestampGranularities []string `json:"timestamp_granularities,omitempty"`
}

type GroqAudioResponse struct {
	Text  string `json:"text"`
	XGroq struct {
		ID string `json:"id"`
	} `json:"x_groq"`
}

func (g GroqTextResponse) GetContent() string {
	var text string
	for _, choice := range g.Choices {
		text += choice.Message.Content
	}
	return text
}

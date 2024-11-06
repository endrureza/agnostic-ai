package models

type ClaudeTextRequest struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
	MaxTokens int `json:"max_tokens"`
	Metadata  *struct {
		UserId *string `json:"user_id"`
	} `json:"metadata,omitempty"`
	StopSequences []string                 `json:"stop_sequences,omitempty"`
	Stream        bool                     `json:"stream,omitempty"`
	System        string                   `json:"system,omitempty"`
	Temperature   float64                  `json:"temperature,omitempty"`
	ToolChoice    map[string]interface{}   `json:"tool_choice,omitempty"`
	Tools         []map[string]interface{} `json:"tools,omitempty"`
	TopP          float64                  `json:"top_p,omitempty"`
	TopK          int                      `json:"top_k,omitempty"`
}

type ClaudeTextResponse struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Role    string `json:"role"`
	Content []struct {
		Text string `json:"text"`
		Type string `json:"type"`
	} `json:"content"`
	Model        string  `json:"model"`
	StopReason   *string `json:"stop_reason"`
	StopSequence *string `json:"stop_sequence"`
	Usage        struct {
		InputTokens  int `json:"input_tokens"`
		OutputTokens int `json:"output_tokens"`
	} `json:"usage"`
}

type ClaudeImageRequest struct{}

type ClaudeImageResponse struct{}

type ClaudeAudioRequest struct{}

type ClaudeAudioResponse struct{}

func (c ClaudeTextResponse) GetContent() string {
	return c.Content[0].Text
}

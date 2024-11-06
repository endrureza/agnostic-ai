package models

type OllamaTextRequest struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
	Stream    bool                     `json:"stream" default:"false"`
	Format    string                   `json:"format,omitempty"`
	Tools     []map[string]interface{} `json:"tools,omitempty"`
	KeepAlive int                      `json:"keep_alive,omitempty"`
	Options   *struct {
		Temperature float64 `json:"temperature,omitempty"`
	} `json:"options,omitempty"`
}
type OllamaTextResponse struct {
	Model     string `json:"model"`
	CreatedAt string `json:"created_at"`
	Message   struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"message,omitempty"`
	DoneReason         string `json:"done_reason,omitempty"`
	Done               bool   `json:"done"`
	TotalDuration      int    `json:"total_duration"`
	LoadDuration       int    `json:"load_duration"`
	PromptEvalCount    int    `json:"prompt_eval_count"`
	PromptEvalDuration int    `json:"prompt_eval_duration"`
	EvalCount          int    `json:"eval_count"`
	EvalDuration       int    `json:"eval_duration"`
}

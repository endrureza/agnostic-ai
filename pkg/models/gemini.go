package models

type GeminiTextRequest struct {
	Contents []struct {
		Role  string `json:"role"`
		Parts []struct {
			Text string `json:"text"`
		} `json:"parts"`
	} `json:"contents"`
	Tools *[]struct {
		FunctionDeclaration *[]struct {
			Name        string                 `json:"name"`
			Description string                 `json:"description"`
			Parameters  map[string]interface{} `json:"parameters,omitempty"`
		} `json:"functionDeclaration,omitempty"`
		GoogleSearchRetrieval *struct {
			DynamicRetrievalConfig interface{} `json:"dynamicRetrievalConfig"`
		} `json:"googleSearchRetrieval,omitempty"`
		CodeExecution *struct {
			FunctionCallingConfig interface{} `json:"functionCallingConfig"`
		} `json:"codeExecution"`
	} `json:"tools,omitempty"`
	ToolConfig *struct {
		FunctionCallingConfig interface{} `json:"functionCallingConfig"`
	} `json:"toolConfig,omitempty"`
	SafetySettings *[]struct {
		Category  string `json:"category"`
		Threshold string `json:"threshold"`
	} `json:"safetySettings,omitempty"`
	GenerationConfig *struct {
		StopSequences    *[]string              `json:"stopSequences,omitempty"`
		ResponseMimeType string                 `json:"responseMimeType,omitempty"`
		ResponseSchema   map[string]interface{} `json:"responseSchema,omitempty"`
		CandidateCount   int                    `json:"candidateCount,omitempty"`
		MaxOutputTokens  int                    `json:"maxOutputTokens,omitempty"`
		Temperature      float64                `json:"temperature,omitempty"`
		TopP             float64                `json:"topP,omitempty"`
		TopK             int                    `json:"topK,omitempty"`
		PresencePenalty  float64                `json:"presencePenalty,omitempty"`
		FrequencyPenalty float64                `json:"frequencyPenalty,omitempty"`
		ResponseLogprobs bool                   `json:"responseLogprobs,omitempty"`
		Logprobs         int                    `json:"logprobs,omitempty"`
	} `json:"generationConfig,omitempty"`
	CachedContent string `json:"cachedContent,omitempty"`
}

type GeminiTextResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
			Role string `json:"role"`
		} `json:"content"`
		FinishReason        string                   `json:"finishReason"`
		SafetyRatings       []map[string]interface{} `json:"safetyRatings"`
		CitationMetadata    map[string]interface{}   `json:"citationMetadata"`
		TokenCount          int                      `json:"tokenCount"`
		GroundingAttributes []map[string]interface{} `json:"groundingAttributes"`
		GroundingMetadata   map[string]interface{}   `json:"groundingMetadata"`
		AvgLogprobs         float64                  `json:"avgLogprobs"`
		LogprobsResult      map[string]interface{}   `json:"logprobsResult"`
		Index               int                      `json:"index"`
	} `json:"candidates"`
}

type GeminiImageRequest struct{}

type GeminiImageResponse struct{}

type GeminiAudioRequest struct{}

type GeminiAudioResponse struct{}

func (g GeminiTextResponse) GetContent() string {
	return g.Candidates[0].Content.Parts[0].Text
}

package qwenmodel

const (
	ChatUser = "user"
	ChatBot  = "assistant"
)

const (
	ChatBaseUrl = "https://dashscope.aliyuncs.com/api/v1/services/aigc/text-generation/generation"
)

type Input struct {
	Messages []Messages `json:"messages"`
}

type QWenTurbo struct {
	Model      string     `json:"model"`
	Input      Input      `json:"input"`
	Parameters Parameters `json:"parameters"`
}

type Parameters struct {
	EnableSearch bool `json:"enable_search"`
}

type Messages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Result

type Output struct {
	Text         string `json:"text"`
	FinishReason string `json:"finish_reason"`
}

type Usage struct {
	OutputTokens int `json:"output_tokens"`
	InputTokens  int `json:"input_tokens"`
}

type Response struct {
	Message   Messages `json:"message"`
	Output    Output   `json:"output"`
	Usage     Usage    `json:"usage"`
	RequestID string   `json:"request_id"`
}

package entity

type OpenaiChatRequest struct {
	Model    string   `json:"model"`
	Messages Messages `json:"messages"`
	Stream   bool     `json:"stream"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Messages []Message

func (msgs *Messages) add(role, content string) {
	*msgs = append(*msgs, Message{
		Role:    role,
		Content: content,
	})
}
func (msgs *Messages) AddUser(content string) {
	msgs.add("user", content)
}

func (msgs *Messages) AddAssistant(content string) {
	msgs.add("assistant", content)
}

type OpenaiChatResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

type OpenaiChatStreamResponse struct {
	Choices []struct {
		Delta struct {
			Content string `json:"content"`
		} `json:"delta"`
	} `json:"choices"`
}

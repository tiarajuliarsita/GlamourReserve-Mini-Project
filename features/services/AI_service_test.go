package services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockOpenAIStub adalah pengganti sederhana untuk OpenAI API dalam pengujian


type ChatMessage struct {
    Role    string `json:"role"`
    Content string `json:"content"`
}

type ChatCompletionRequest struct {
    Model    string           `json:"model"`
    Messages []ChatMessage    `json:"messages"`
}

type ChatChoice struct {
    Message ChatMessage `json:"message"`
}

type ChatCompletionResponse struct {
    Choices []ChatChoice `json:"choices"`
}
type MockOpenAIStub struct{}

func (s *MockOpenAIStub) CreateChatCompletion(ctx context.Context, req ChatCompletionRequest) (ChatCompletionResponse, error) {
	// Implementasi palsu API yang selaras dengan kebutuhan pengujian Anda
	response := ChatCompletionResponse{
		Choices: []ChatChoice{
			{
				Message: ChatMessage{
					Content: "Jawaban dari OpenAI",
				},
			},
		},
	}
	return response, nil
}

func TestBeautyCareService_AskAboutBeauty(t *testing.T) {

	beautyCareService := NewBeautyCare()
	userInput := "Apa perawatan kecantikan yang cocok untuk saya?"
	brand := "Glam"
	skinType := "Normal"

	// Menjalankan pengujian
	beautyCareService.AskAboutBeauty(userInput, brand, skinType, "YOUR_OPENAI_API_KEY")
	assert.Equal(t, "", "")
}

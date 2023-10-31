package services

import (
	"context"
	"glamour_reserve/entity/request"
	"os"

	"github.com/sashabaranov/go-openai"
)

type beautyCareService struct {
}

type BeautyCareInterface interface {
	AskAboutBeauty(question request.AskBeautyReq, openAIKey string) (string, error)
}

func NewBeautyCare() BeautyCareInterface {
	return &beautyCareService{}
}

func (s *beautyCareService) AskAboutBeauty(question request.AskBeautyReq, openAIKey string) (string, error) {
	ctx := context.Background()
	client := openai.NewClient(openAIKey)
	model := openai.GPT3Dot5Turbo
	filePath := "utils/helpers/prompt/prompt.txt"
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: string(content),
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: question.Question,
		},
	}

	resp, err := s.getCompletionFromMessages(ctx, client, messages, model)
	if err != nil {
		return "", err
	}
	answer := resp.Choices[0].Message.Content
	return answer, nil
}

func (s *beautyCareService) getCompletionFromMessages(
	ctx context.Context,
	client *openai.Client,
	messages []openai.ChatCompletionMessage,
	model string,
) (openai.ChatCompletionResponse, error) {
	if model == "" {
		model = openai.GPT3Dot5Turbo
	}

	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: messages,
		},
	)
	return resp, err

}

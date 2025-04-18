package ai

import (
	"context"
	"github.com/sashabaranov/go-openai"
)

// Функция запроса к OpenAI
func GetChatResponse(client *openai.Client, input string) (string, error) {

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    "user",
					Content: input,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

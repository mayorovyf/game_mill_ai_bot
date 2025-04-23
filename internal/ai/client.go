// internal/ai/client.go
package ai

import (
	"context"
	"github.com/sashabaranov/go-openai"
)

// временная реализация
// Функция запроса к OpenAI
func GetChatResponse(client *openai.Client, input string) (string, error) {

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: "gpt-4.1-nano-2025-04-14",
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

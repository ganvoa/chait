package internal

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/sashabaranov/go-openai"
)

type AIParticipant struct {
	Name     string
	Role     string
	client   *openai.Client
	messages []string
	logger   logr.Logger
}

func NewAiParticipant(name string, role string, apiKey string, logger logr.Logger) *AIParticipant {
	var messages []string

	client := openai.NewClient(apiKey)

	return &AIParticipant{
		Name:     name,
		client:   client,
		messages: messages,
		Role:     role,
		logger:   logger,
	}
}

func (p *AIParticipant) Talk() (*ChatMessage, error) {

	p.logger.Info("calling Talk CreateChatCompletion")
	resp, err := p.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: p.Role,
				},
			},
		},
	)

	if err != nil {
		return nil, err
	}

	message := resp.Choices[0].Message.Content

	p.logger.Info("message received Talk CreateChatCompletion", "total_tokens", resp.Usage.TotalTokens, "finish_reason", resp.Choices[0].FinishReason)
	return &ChatMessage{p.Name, message}, nil
}

func (p *AIParticipant) Reply(cm *ChatMessage) (*ChatMessage, error) {
	p.logger.Info("calling Reply CreateChatCompletion")
	resp, err := p.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: p.Role,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: cm.Message,
				},
			},
		},
	)

	if err != nil {
		return nil, err
	}

	message := resp.Choices[0].Message.Content
	p.logger.Info("message received Talk CreateChatCompletion", "total_tokens", resp.Usage.TotalTokens, "finish_reason", resp.Choices[0].FinishReason)

	return &ChatMessage{p.Name, message}, nil
}

package ai

import (
	"fmt"

	"github.com/leehosu/commitmate/internal/config"
)

// GenerateOptions는 커밋 메시지 생성 옵션입니다
type GenerateOptions struct {
	Temperature     float64 // AI 생성 다양성 (0.0 ~ 2.0)
	PreviousMessage string  // 이전에 생성된 메시지 (재생성 시 사용)
}

// DefaultGenerateOptions는 기본 생성 옵션을 반환합니다
func DefaultGenerateOptions() GenerateOptions {
	return GenerateOptions{
		Temperature:     0.7,
		PreviousMessage: "",
	}
}

// RegenerateOptions는 재생성 시 사용할 옵션을 반환합니다
func RegenerateOptions(previousMessage string) GenerateOptions {
	return GenerateOptions{
		Temperature:     1.0, // 재생성 시 더 높은 다양성
		PreviousMessage: previousMessage,
	}
}

// Client는 AI 클라이언트 인터페이스입니다
type Client interface {
	GenerateCommitMessage(systemPrompt, userPrompt string, opts GenerateOptions) (string, error)
}

// NewClient는 설정에 따라 적절한 AI 클라이언트를 생성합니다
func NewClient(cfg *config.Config) (Client, error) {
	switch cfg.Provider {
	case "openai":
		if cfg.OpenAI.APIKey == "" {
			return nil, fmt.Errorf("OpenAI API 키가 설정되지 않았습니다")
		}
		return NewOpenAIClient(cfg.OpenAI.APIKey, cfg.OpenAI.Model, cfg.OpenAI.MaxTokens), nil
	
	case "claude":
		if cfg.Claude.APIKey == "" {
			return nil, fmt.Errorf("Claude API 키가 설정되지 않았습니다")
		}
		return NewClaudeClient(cfg.Claude.APIKey, cfg.Claude.Model, cfg.Claude.MaxTokens), nil
	
	default:
		return nil, fmt.Errorf("지원하지 않는 제공자입니다: %s", cfg.Provider)
	}
}

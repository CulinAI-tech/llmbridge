package llm

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/genai"
)

type LlmService struct {
	provider string
	apiKey   string
	url      string
}

func NewLlmService(provider, apiKey, url string) *LlmService {
	return &LlmService{
		provider: provider,
		apiKey:   apiKey,
		url:      url,
	}
}

func (s *LlmService) Query(ctx context.Context, prompt string) (*QueryResponse, error) {
	switch s.provider {
	case "gemini":
		return s.queryGemini(ctx, prompt)
	default:
		return nil, ErrUnsupportedProvider
	}
}

func (s *LlmService) queryGemini(ctx context.Context, prompt string) (*QueryResponse, error) {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: s.apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, err
	}
	response, err := client.Models.GenerateContent(
		ctx,
        "gemini-2.5-flash-lite",
        genai.Text(prompt),
        nil,
	)
	if err  != nil {
		return nil, fmt.Errorf("Gemini API error: %w", err)
	}
	
	fmt.Println("Gemini response:", response.Text())

	return &QueryResponse{
		Answer:    response.Text(),
		Query:     prompt,
		Timestamp: time.Now(),		
	}, nil	
}
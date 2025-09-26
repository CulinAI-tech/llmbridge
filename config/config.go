package config

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	LLMApiKey string `env:"LLM_API_KEY" envDefault:""`
	LLMProvider string `env:"LLM_PROVIDER" envDefault:"openai"`
	LLMUrl string `env:"LLM_URL" envDefault:"https://api.openai.com/v1/chat/completions"`
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("failed to load config: %v", err)
		return nil, err
	}
	return &cfg, nil
}
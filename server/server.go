package server

import (
	"llmbridge/config"
	"llmbridge/llm"
	"llmbridge/middleware"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	cfg *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		router: gin.Default(),
		cfg: cfg,
	}
}

func (s *Server) Setup() error {
	s.router.Use(gin.Logger())
	s.router.Use(gin.Recovery())
	s.router.Use(middleware.SetupCORS())

	

	llmService := llm.NewLlmService(s.cfg.LLMProvider, s.cfg.LLMApiKey, s.cfg.LLMUrl)
	llmHandler := llm.NewHandler(llmService)

	api := s.router.Group("/api")
	api.POST("/llm/query", llmHandler.Query)

	return nil
}

func (s *Server) Run() error {
	return s.router.Run(":8080")
}
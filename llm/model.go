package llm

import "time"

type QueryResponse struct {
	Answer    string      `json:"answer"`
	Query     string      `json:"query,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Timestamp time.Time   `json:"timestamp"`
}

type QueryRequest struct {
    Prompt string `json:"prompt"`
}
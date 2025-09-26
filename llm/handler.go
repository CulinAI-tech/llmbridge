package llm

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LlmHandler struct {
	service *LlmService
}

func NewHandler(service *LlmService) *LlmHandler {
	return &LlmHandler{
		service: service,
	}
}

func (h *LlmHandler) Query(c *gin.Context) {
	var req QueryRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := h.service.Query(c.Request.Context(), req.Prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
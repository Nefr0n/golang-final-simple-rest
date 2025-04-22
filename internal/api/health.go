package api

import (
	"github.com/Skapar/simple-rest/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// CheckHealthHandler godoc
// @Summary Check the health of the server
// @Description Check if the server is running
// @Tags health
// @Accept  json
// @Produce  json
// @Success 200 {string} string
// @Router /health [get]
func (hh *HealthHandler) CheckHealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, models.Response[interface{}]{
		Data:    nil,
		Message: "server is running",
		Success: true,
	})
}

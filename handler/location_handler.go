package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LocationHandler struct{}

func NewLocationHandler() *LocationHandler {
	return &LocationHandler{}
}

func (h *LocationHandler) GetLocations(c *gin.Context) {
	c.JSON(http.StatusOK, []any{})
}

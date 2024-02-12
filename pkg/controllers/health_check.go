package controllers

import (
	"github.com/gin-gonic/gin"
)

type HealthCheckController struct{}

func (ct HealthCheckController) Index(c *gin.Context) {
	c.JSON(200, map[string]string{"status": "OK"})
}

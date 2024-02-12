package controllers

import (
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	healthCtl := HealthCheckController{}
	r.GET("/healthcheck", healthCtl.Index)

	return r
}

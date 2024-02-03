package controllers

import (
	"net/http"

	"github.com/diskymg/go-challenge/diskymg/models"
)

func initialize() http.Handler {
	if err := models.InitDB("localhost", "postgres", "koikeya", "ocpi_test", "5432", true, true); err != nil {
		panic("failed to InitDB")
	}

	ocpiApi := NewOcpiApi()
	ginPetServer := NewGinOcpiServer(ocpiApi, "8080")
	return ginPetServer.Handler
}

func finalize() {
	models.CloseDB()
	CloseGinOcpiServer()
}

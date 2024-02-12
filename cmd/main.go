package main

import (
	"go-challenge/pkg/config"
	db "go-challenge/pkg/database"
	"go-challenge/pkg/router"
)

func main() {
	config.Init()
	db.Init()
	defer db.Close()
	router.Init()
}

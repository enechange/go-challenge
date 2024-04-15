package main

import (
	"go-challenge/configs"
	db "go-challenge/internal/infrastructure/database"
	"log"
)

func main() {
	configs.Init()
	db.Init()
	err := db.MigrateExecution()
	if err != nil {
		log.Print(err)
	}
	defer db.Close()
	serverInit()
}

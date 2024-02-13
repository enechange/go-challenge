package main

import (
	"go-challenge/config"
	db "go-challenge/database"
)

func main() {
	config.Init()
	db.Init()
	defer db.Close()
	serverInit()
}

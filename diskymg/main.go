package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/diskymg/go-challenge/diskymg/controllers"
	"github.com/diskymg/go-challenge/diskymg/models"
)

func main() {
	appPort := flag.String("port", os.Getenv("APP_PORT"), "Port for test HTTP server")
	flag.Parse()

	// database
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_DATABASE")
	port := os.Getenv("DB_PORT")

	if err := models.InitDB(host, user, password, dbname, port, true, true); err != nil {
		fmt.Fprintf(os.Stderr, "Error InitDB\n: %s", err)
	}
	defer models.CloseDB()

	// Create an instance of our handler which satisfies the generated interface
	ocpiApi := controllers.NewOcpiApi()
	s := controllers.NewGinOcpiServer(ocpiApi, *appPort)
	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}

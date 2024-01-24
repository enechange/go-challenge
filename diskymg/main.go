package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"

	"github.com/diskymg/go-challenge/diskymg/controllers"
	"github.com/diskymg/go-challenge/diskymg/models"
	"github.com/diskymg/go-challenge/diskymg/ocpi"
)

func NewGinOcpiServer(ocpiApi *controllers.OcpiApi, port string) *http.Server {
	swagger, err := ocpi.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// This is how you set up a basic gin router
	r := gin.Default()

	// CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8001"}
	r.Use(cors.New(corsConfig))

	// BaseURL
	baseURL := "/ocpi/diskymg/2.2"

	// /openapi-spec.yaml endpoint
	r.GET(baseURL+"/openapi-spec.yaml", func(c *gin.Context) { c.File("ocpi/openapi-spec.yaml") })

	// We now register our ocpiApi above as the handler for the interface
	ocpi.RegisterHandlersWithOptions(r, ocpiApi, ocpi.GinServerOptions{BaseURL: baseURL})

	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	r.Use(middleware.OapiRequestValidator(swagger))

	s := &http.Server{
		Handler: r,
		Addr:    net.JoinHostPort("0.0.0.0", port),
	}
	return s
}

func main() {
	port := flag.String("port", os.Getenv("APP_PORT"), "Port for test HTTP server")
	flag.Parse()

	// database
	if err := models.InitDB(); err != nil {
		fmt.Fprintf(os.Stderr, "Error InitDB\n: %s", err)
	}
	if err := models.AutoMigrate(); err != nil {
		fmt.Fprintf(os.Stderr, "Error AutoMigrate\n: %s", err)
	}
	if err := models.Seed(); err != nil {
		fmt.Fprintf(os.Stderr, "Error Seed\n: %s", err)
	}
	defer models.CloseDB()

	// Create an instance of our handler which satisfies the generated interface
	ocpiApi := controllers.NewOcpiApi()
	s := NewGinOcpiServer(ocpiApi, *port)
	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}

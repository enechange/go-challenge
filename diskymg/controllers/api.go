package controllers

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"

	"github.com/diskymg/go-challenge/diskymg/ocpi"
)

var s *http.Server

type OcpiApi struct {
}

func NewOcpiApi() *OcpiApi {
	return &OcpiApi{}
}

func NewGinOcpiServer(ocpiApi *OcpiApi, port string) *http.Server {
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

	s = &http.Server{
		Handler: r,
		Addr:    net.JoinHostPort("0.0.0.0", port),
	}
	return s
}

func CloseGinOcpiServer() {
	s.Close()
}

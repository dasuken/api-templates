package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/noguchidaisuke/api-templates/handlers"
	"github.com/noguchidaisuke/api-templates/routes"
	"log"
	"net/http"
)

var (
	port     int
)

func init() {
	flag.IntVar(&port, "port", 9000, "api service port")
	flag.Parse()
}

func main() {
	// create handlers
	authHandlers := handlers.NewAuthHandlers()

	// create routes instance
	healthRoutes := routes.NewHealthRoutes()
	authRoutes := routes.NewAuthRoutes(authHandlers)

	// install for routers
	router := mux.NewRouter().StrictSlash(true)
	routes.Install(router, healthRoutes)
	routes.Install(router, authRoutes)

	log.Printf("API service runnning on [::]:%d\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), routes.WithCORS(router)))
}
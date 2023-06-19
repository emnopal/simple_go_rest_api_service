package main

import (
	"log"

	"github.com/emnopal/simple_go_rest_api_service/pkg/config"
	"github.com/emnopal/simple_go_rest_api_service/pkg/routes"
)

func main() {
	log.Println("Listing for requests at http://localhost:8000/")
	routes.Routes()
	config.ServerConfig()
}

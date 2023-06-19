package config

import (
	"log"
	"net/http"

	helper "github.com/emnopal/go_helper"
)

func ServerConfig() {
	PORTS := helper.GetENV("PORTS", ":8000")
	log.Fatal(http.ListenAndServe(PORTS, nil))
}

package routes

import (
	"net/http"

	controller "github.com/emnopal/simple_go_rest_api_service/pkg/controllers"
)

func Routes() {
	http.HandleFunc("/", controller.HelloWorld)
}

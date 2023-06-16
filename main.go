package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	helper "github.com/emnopal/go_helper"
)

type server struct{}

type ExampleJSON struct {
	JSONBody string `json:"json_body,omitempty"`
}

func (s *server) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	// to set which origin can access this rest api
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// to set which methods is allowed to access this rest api
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")

	// to set which headers is allowed to access this rest api
	w.Header().Set("Access-Control-Allow-Headers", "*")

	// added headers
	w.Header().Set("Content-Type", "application/json")

	message := ""
	status := http.StatusOK

	query_param := req.URL.Query().Get("query_param")

	JSONBody := ""

	var t ExampleJSON
	if err := json.NewDecoder(req.Body).Decode(&t); err != nil && req.Method == "POST" {
		log.Print("WARNING! JSON is empty")
	} else {
		JSONBody = t.JSONBody
	}

	switch req.Method {
	case "GET":
		message = fmt.Sprintf(`{
			"message": "Hello World from GET",
			"query_params": "%s"
		}`, query_param)
	case "POST":
		message = fmt.Sprintf(`{
			"message": "Hello World from POST",
			"json_body": "%s"
		}`, JSONBody)
	default:
		message = fmt.Sprintf(`{"message": "Method %s not allowed"}`, req.Method)
		status = http.StatusMethodNotAllowed
	}

	w.WriteHeader(status)
	w.Write([]byte(message))

}

func main() {
	log.Println("Listing for requests at http://localhost:8000/")
	PORTS := helper.GetENV("PORTS", ":8000")
	serve := &server{}
	http.HandleFunc("/", serve.ServeHTTP)
	log.Fatal(http.ListenAndServe(PORTS, nil))
}

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

	headParams := &helper.HeaderParams{
		AccessControlAllowMethods: "GET, POST",
	}
	helper.SetHeader(w, headParams)

	message := ""
	status := http.StatusOK

	query_param := req.URL.Query().Get("query_param")

	var t ExampleJSON
	JSONBody := ""

	if req.Method == "POST" {
		err := json.NewDecoder(req.Body).Decode(&t)
		if err != nil {
			log.Print("WARNING! JSON is empty")
		} else {
			JSONBody = t.JSONBody
		}
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

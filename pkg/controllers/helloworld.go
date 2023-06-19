package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	helper "github.com/emnopal/go_helper"
	schemas "github.com/emnopal/simple_go_rest_api_service/pkg/schemas/json"
)

func HelloWorld(w http.ResponseWriter, req *http.Request) {

	headParams := &helper.HeaderParams{
		AccessControlAllowMethods: "GET, POST",
	}
	helper.SetHeader(w, headParams)

	message := ""
	status := http.StatusOK

	query_param := req.URL.Query().Get("query_param")

	var t schemas.ExampleJSON
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

package main

import (
	"io"
	"log"
	"net/http"

	helper "github.com/emnopal/go_helper"
)

func main() {

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello World!")
	}

	http.HandleFunc("/hello", helloHandler)

	log.Println("Listing for requests at http://localhost:8000/hello")

	PORTS := helper.GetENV("PORTS", ":8000")

	log.Fatal(http.ListenAndServe(PORTS, nil))

}

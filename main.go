package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func GetENV(key string, default_value string) (ENV_VALUE string) {
	ENV_VALUE = os.Getenv(key)
	if ENV_VALUE == "" {
		ENV_VALUE = default_value
	}
	return
}

func main() {

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello World!")
	}

	http.HandleFunc("/hello", helloHandler)

	log.Println("Listing for requests at http://localhost:8000/hello")

	PORTS := GetENV("PORTS", ":8000")

	log.Fatal(http.ListenAndServe(PORTS, nil))

}

package main

import (
	"log"
	"net/http"

	"github.com/saireddyb/kubeide/pkg/handler"
)

func main() {
	apiHandler, err := handler.CreateHTTPAPIHandler()
	if err != nil {
		log.Fatalf("error creating API handler: %v", err)
	}
	http.Handle("/", apiHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// API microservice for logging in and storing trudie dos.
package main

import (
	"log"
	"net/http"

	"github.com/vallard/trudiedo-api/routes"
)

func main() {
	router := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

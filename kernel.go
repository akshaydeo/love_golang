package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func getPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "4747"
	}
	return ":" + port
}
func main() {
	log.Println("Starting main server")
	r := mux.NewRouter()
	pushRoutes(r)
	http.Handle("/", r)
	http.ListenAndServe(getPort(), nil)
}

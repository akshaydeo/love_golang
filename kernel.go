package main

import (
	"github.com/gorilla/mux"
	"log"
	"love_golang/dbHelper"
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
	log.Println("Initializing database...")
	if err := dbHelper.IntializeDatabase(); err != nil {
		panic("Error while initializing database")
		return
	}
	log.Println("Database initialized...")
	r := mux.NewRouter()
	pushRoutes(r)
	http.Handle("/", r)
	// setting up public assets path
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	http.ListenAndServe(getPort(), nil)
}

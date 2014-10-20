package main

import (
	"github.com/gorilla/mux"
	"love_golang/controllers"
)

func pushRoutes(r *mux.Router) {
	r.HandleFunc("/", controllers.Home).Methods("GET")
}

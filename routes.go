package main

import (
	"github.com/gorilla/mux"
	"love_golang/controllers"
)

func pushRoutes(r *mux.Router) {
	r.HandleFunc("/", controllers.Home).Methods("GET")

	r.HandleFunc("/user/login", controllers.LoginWithTwitter).Methods("POST")

	r.HandleFunc("/post", controllers.CreatePost).Methods("POST")
	r.HandleFunc("/post/{post_id}/upvote", controllers.CreatePost).Methods("POST")
}

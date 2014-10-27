package controllers

import (
	"log"
	"love_golang/models"
	"net/http"
)

func CreatePost(rw http.ResponseWriter, req *http.Request) {
	log.Println("Creating a post", req.FormValue("post"))
	if req.FormValue("post") == "" {
		http.Redirect(rw, req, "/", http.StatusSeeOther)
		return
	}
	twitterHandle := req.FormValue("twitter_handle")
	if twitterHandle == "" {
		twitterHandle = "@anonymous"
	}
	user := models.User{TwitterHandle: twitterHandle}
	user.Save()
	log.Println("User id is", user.Id)
	post := models.Post{Post: req.FormValue("post"), PostedByUserId: user.Id}
	post.Save()
	http.Redirect(rw, req, "/", http.StatusSeeOther)
}

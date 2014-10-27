package controllers

import (
	"github.com/yosssi/ace"
	"love_golang/models"
	"net/http"
)

func Home(rw http.ResponseWriter, req *http.Request) {
	tpl, err := ace.Load("views/home", "", &ace.Options{DynamicReload: true})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	posts, err := models.GetAllPosts()
	if err != nil {
		if err := tpl.Execute(rw, nil); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	data := map[string]interface{}{
		"Posts": posts,
	}
	if err := tpl.Execute(rw, data); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

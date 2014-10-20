package controllers

import (
	"github.com/yosssi/ace"
	"net/http"
)

func Home(rw http.ResponseWriter, req *http.Request) {
	tpl, err := ace.Load("views/home", "", &ace.Options{DynamicReload: true})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	if err := tpl.Execute(rw, nil); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

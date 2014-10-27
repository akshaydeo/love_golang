package controllers

import (
	"github.com/ChimeraCoder/anaconda"
	"log"
	"net/http"
	"os"
)

func LoginWithTwitter(rw http.ResponseWriter, req *http.Request) {
	log.Println("Logging in with twitter")
	anaconda.SetConsumerKey(os.Getenv("TWITTER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTER_SECRET"))
	url, _, err := anaconda.AuthorizationURL("/")
	log.Println("executed authorization")
	if err != nil {
		log.Println(err.Error())
		http.Redirect(rw, req, "/", 200)
		return
	}
	log.Println("Got url as", url)
	http.Redirect(rw, req, url, 200)
}

func AuthorizationCallback(rw http.ResponseWriter, req *http.Request) {
	log.Println("Got authorization callback", req.RequestURI)
}

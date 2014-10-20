package controllers

import (
	"fmt"
	"net/http"
)

func Home(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "Testing page")
}

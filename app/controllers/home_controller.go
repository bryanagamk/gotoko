package controllers

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf("Welcome to GoToko Home Page")
	fmt.Fprintf(w, "Welcome to GoToko Homepage")
}

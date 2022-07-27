package controllers

import "net/http"

func signUp(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "signup")
}

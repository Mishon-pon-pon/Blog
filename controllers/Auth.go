package controllers

import (
	"html/template"
	"net/http"
)

//Auth - handler for contact page
func Auth(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles(
		"web/auth.html",
	))
	tmp.Execute(w, nil)
}

package controllers

import (
	"html/template"
	"net/http"
)

//About - handler for contact page
func About(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles(
		"web/about.html",
	))
	tmp.Execute(w, nil)
}

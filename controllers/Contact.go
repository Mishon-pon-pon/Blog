package controllers

import (
	"html/template"
	"net/http"
)

//Contact - handler for contact page
func Contact(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles(
		"web/contact.html",
	))
	tmp.Execute(w, nil)
}

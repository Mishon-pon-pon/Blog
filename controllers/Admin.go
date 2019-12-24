package controllers

import (
	"html/template"
	"net/http"
)

//Admin - handler /admin
func Admin(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles(
		"web/Admin.html",
	))
	tmp.Execute(w, nil)
}

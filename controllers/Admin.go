package controllers

import (
	"Blog/libs/dataBase"
	"Blog/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

//Admin - handler /admin
func Admin(w http.ResponseWriter, r *http.Request) {
	e := r.URL.Query().Get("email")
	p := r.URL.Query().Get("password")
	dataBase.DBQuery(`Select * from Users where Email = '`+e+`' and Pass = '`+p+`'`, func(result *sql.Rows) {
		u := models.User{}
		for result.Next() {
			err := result.Scan(&u.UserID, &u.Email, &u.Pass)
			if err != nil {
				panic(err)
			}
		}
		if u.UserID != 0 {
			tmp := template.Must(template.ParseFiles(
				"web/Admin.html",
			))
			tmp.Execute(w, nil)
		} else {
			http.Redirect(w, r, "/auth", 301)
		}
	})

}

//NewPost - write new post into database
func NewPost(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	var a models.Article
	err := dec.Decode(&a)
	if err != nil {
		fmt.Println(err)
	}
	dataBase.DBExec(`Insert into Articles(Title, TextArticle) Values('` + a.Title + `', '` + a.TextArticle + `')`)
}

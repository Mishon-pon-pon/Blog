package controllers

import (
	"Blog/libs/dataBase"
	"Blog/models"
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
)

//SingleArticle -
func SingleArticle(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	tmp := template.Must(template.ParseFiles(
		"web/single.html",
	))
	a := models.Article{}
	dataBase.DBQuery(`Select * from Articles where ArticleId = `+id, func(result *sql.Rows) {

		for result.Next() {
			err := result.Scan(&a.ID, &a.Title, &a.TextArticle)
			if err != nil {
				fmt.Println(err)
			}
		}
		tmp.Execute(w, a)
	})

}

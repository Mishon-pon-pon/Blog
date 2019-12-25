package controllers

import (
	"Blog/libs/dataBase"
	"Blog/models"
	"database/sql"
	"html/template"
	"net/http"
)

//IndexHandler - handler for index page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	index := template.Must(template.ParseFiles(
		"web/index.html",
	))
	articles := []models.Article{}
	dataBase.DBQuery(`SELECT 
						ArticleId, 
						Title, 
						substr(TextArticle, 0, 219) || ltrim(substr(TextArticle, 219, 220), ' ') || '...' as TextArticle   
						FROM Articles`, func(result *sql.Rows) {
		a := models.Article{}
		for result.Next() {
			err := result.Scan(&a.ID, &a.Title, &a.TextArticle)
			if err != nil {
				panic(err)
			}
			articles = append(articles, a)
		}
	})
	index.Execute(w, articles)
}

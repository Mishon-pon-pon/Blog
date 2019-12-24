package main

import (
	"Blog/controllers"
	"Blog/libs/dataBase"
	"Blog/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	oneMux := http.NewServeMux()

	oneMux.HandleFunc("/", indexHandler)
	oneMux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web"))))
	oneMux.HandleFunc("/test", controllers.PersonCreate)
	oneMux.HandleFunc("/admin", controllers.Admin)
	oneMux.HandleFunc("/article", controllers.SingleArticle)
	fmt.Println("server run on 3003")
	http.ListenAndServe(":3003", oneMux)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	index := template.Must(template.ParseFiles(
		"web/index.html",
	))
	articles := []models.Article{}
	dataBase.DBQuery(`SELECT ArticleId, Title, substr(TextArticle, 0, 220) as TextArticle   FROM Articles`, func(result *sql.Rows) {
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

func newPost(w http.ResponseWriter, r *http.Request) {
	var test struct {
		test string
	}
	err := json.NewDecoder(r.Body).Decode(&test)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(test)
}

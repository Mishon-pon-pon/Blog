package main

import (
	"Blog/controllers"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	mainMux := http.NewServeMux()

	mainMux.HandleFunc("/", controllers.IndexHandler)
	mainMux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web"))))
	mainMux.HandleFunc("/admin", controllers.Admin)
	mainMux.HandleFunc("/new", controllers.NewPost)
	mainMux.HandleFunc("/article", controllers.SingleArticle)
	fmt.Println("server run on 3003")
	http.ListenAndServe(":3003", mainMux)
}

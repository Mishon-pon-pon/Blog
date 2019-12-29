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
	mainMux.HandleFunc("/contact", controllers.Contact)
	mainMux.HandleFunc("/about", controllers.About)
	mainMux.HandleFunc("/auth", controllers.Auth)
	mainMux.HandleFunc("/userauth", func(w http.ResponseWriter, r *http.Request) {
		n := r.URL.Query().Get("email")
		fmt.Println(n)
	})
	// http.Get("http://localhost:3003/auth?email=123@123&password=qwe")
	fmt.Println("server run on 3003")
	http.ListenAndServe(":133", mainMux)
}

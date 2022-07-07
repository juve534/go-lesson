package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/juve534/go-lesson/app/models"
	"html"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Get("/bar", bar)
	r.Get("/posts/{id}", postIndex)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func bar(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hello, %q", html.EscapeString(r.URL.Path))))
}

func postIndex(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "id")
	db := models.NewDB()
	post, err := db.GetPostById(postID)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(post); err != nil {
		panic(err)
	}
}

package main

import (
    "database/sql"
    "encoding/json"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
    "github.com/go-chi/chi"
    "html"
    "log"
    "net/http"
)

type Posts struct {
    Id int          `json:"id"`
    Title string    `json:"title"`
    Body string     `json:"body"`
}

func main() {
    r := chi.NewRouter()

    r.Get("/bar", bar)
    r.Get("/posts/{id}", getAll)

    log.Fatal(http.ListenAndServe(":8080", r))
}

func bar(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(fmt.Sprintf("Hello, %q", html.EscapeString(r.URL.Path))))
}

func getAll(w http.ResponseWriter, r *http.Request)  {
    postID := chi.URLParam(r, "id")

    // MySQLと接続
    db, err := sql.Open("mysql", "root:pass@tcp(masterdb:3306)/lesson?parseTime=true")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    // データ取得を実施
    rows, err := db.Query("SELECT id,title,body FROM posts WHERE id=?", postID)
    if err != nil {
        panic(err.Error())
    }

    var id int
    var title string
    var body string

    for rows.Next() {
        if err := rows.Scan(&id, &title, &body); err != nil {
            log.Fatal(err)
        }
        log.Println(id, title, body)
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(Posts{Id: id,Title: body, Body: body}); err != nil {
        panic(err)
    }
}
package main

import (
    "fmt"
    "github.com/go-chi/chi"
    "html"
    "log"
    "net/http"
)

func main() {
    r := chi.NewRouter()

    r.Get("/bar", bar)

    log.Fatal(http.ListenAndServe(":8080", r))
}

func bar(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(fmt.Sprintf("Hello, %q", html.EscapeString(r.URL.Path))))
}
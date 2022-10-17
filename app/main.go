package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/juve534/go-lesson/app/models"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	err := server()
	panic(err)
}

func server() error {
	r := chi.NewRouter()

	r.Get("/bar", bar)

	db := models.NewDB()
	h := NewPostHandler(db)
	r.Get("/posts/{id}", h.postIndex)
	r.Post("/posts/", h.postCreate)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	startServerErr := make(chan error, 1)

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			startServerErr <- err
			db.Close()
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	<-quit
	log.Println("run server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(fmt.Sprintf("failed to graceful shutdown, err = %s", err.Error()))
		return err
	}
	log.Println("successfully graceful shutdown server")
	return nil
}

func bar(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hello, %q", html.EscapeString(r.URL.Path))))
}

type postHandler struct {
	d *models.DB
}

func NewPostHandler(d *models.DB) *postHandler {
	return &postHandler{d: d}
}

func (h postHandler) postIndex(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "id")
	post, err := h.d.GetPostById(postID)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(post); err != nil {
		panic(err)
	}
}

type PostsCreateRequest struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (h postHandler) postCreate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var req PostsCreateRequest
	if err := json.Unmarshal(body, &req); err != nil {
		panic(err)
	}

	post, err := models.NewPosts(req.ID, req.Title, req.Body)
	if err != nil {
		panic(err)
	}

	err = h.d.CreatePost(post)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(post); err != nil {
		panic(err)
	}
}

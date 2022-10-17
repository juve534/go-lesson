package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/juve534/go-lesson/app/handler"
	"github.com/juve534/go-lesson/app/models"
	"html"
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
	model := models.NewPostsModel(db)
	h := handler.NewPostHandler(model)
	r.Get("/posts/{id}", h.PostIndex)
	r.Post("/posts/", h.PostCreate)

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

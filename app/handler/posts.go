package handler

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/juve534/go-lesson/app/models"
	"io/ioutil"
	"net/http"
)

type postHandler struct {
	d models.PostsModel
}

func NewPostHandler(d models.PostsModel) *postHandler {
	return &postHandler{d: d}
}

func (h *postHandler) PostIndex(w http.ResponseWriter, r *http.Request) {
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

func (h *postHandler) PostCreate(w http.ResponseWriter, r *http.Request) {
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

package models

import (
	"database/sql"
	"fmt"
)

type DB struct {
	c *sql.DB
}

func NewDB() *DB {
	db, err := sql.Open("mysql", "root:pass@tcp(masterdb:3306)/lesson?parseTime=true")
	if err != nil {
		panic(err)
	}

	return &DB{
		c: db,
	}
}

type Posts struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func NewPosts(id int, title, body string) (*Posts, error) {
	if title == "" {
		return nil, fmt.Errorf("invalid title")
	}
	if body == "" {
		return nil, fmt.Errorf("invalid body")
	}

	return &Posts{
		ID:    id,
		Title: title,
		Body:  body,
	}, nil
}

func (d *DB) GetPostById(postID string) (*Posts, error) {
	rows, err := d.c.Query("SELECT id,title,body FROM posts WHERE id=?", postID)
	if err != nil {
		return nil, err
	}
	defer d.c.Close()

	var id int
	var title string
	var body string

	for rows.Next() {
		if err := rows.Scan(&id, &title, &body); err != nil {
			return nil, err
		}
	}

	p, err := NewPosts(id, title, body)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (d *DB) CreatePost(posts *Posts) error {
	r, err := d.c.Exec("INSERT INTO posts (title, body, created) VALUES (?, ?, NOW())", posts.Title, posts.Body)
	if err != nil {
		return err
	}
	defer d.c.Close()

	id, err := r.LastInsertId()
	if err != nil {
		return err
	}
	posts.ID = int(id)

	return nil
}

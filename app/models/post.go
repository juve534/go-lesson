package models

import (
	"database/sql"
)

type Db struct {
	c *sql.DB
}

func NewDB() *Db {
	db, err := sql.Open("mysql", "root:pass@tcp(masterdb:3306)/lesson?parseTime=true")
	if err != nil {
		panic(err)
	}

	return &Db{
		c: db,
	}
}

type Posts struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (d *Db) GetPostById(postID string) (*Posts, error) {
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

	return &Posts{
		Id:    id,
		Title: title,
		Body:  body,
	}, nil
}

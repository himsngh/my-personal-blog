package database

import (
	"database/sql"
	"time"
)

type Store interface {
	CreatePost(post *Post) (*Post, error)
	GetPost(id int) (sql.Row, error)
	ListPost() (sql.Rows, error)
	DeletePost(id int) error
	SearchPost(params []string) (sql.Rows, error)
}

type Post struct {
	Id     string
	Title  string
	Data   string
	Date   time.Time
	Author *Author
}

type Author struct {
	Id             string
	FirstName      string
	LastName       string
	Email          string
	ProfilePicture string
}

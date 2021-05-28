package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
	"time"
)

type server struct {
	db *sql.DB
}

func NewServer() (*server, error) {

	database, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		return nil, err
	}
	if err := database.Ping(); err != nil {
		return nil, err
	}

	return &server{
		db:database,
	}, nil
}

func (s *server) Create(post *Post) (*Post, error) {

	query := `
	INSERT INTO personal_blog.v1.post (title, content, created_on)
	VALUES ($1, $2, $3)
`
	res, err := s.db.Exec(query, post.Title, post.Data, time.Now())
	if err != nil {
		return nil, err
	}
	id, err  := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	post.Id = strconv.FormatInt(id, 10)


	return post, nil
}


package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strconv"
	"time"
)

type server struct {
	db *sql.DB
}

func NewDatabaseStore() (Store, error) {

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

func (s *server) CreatePost(post *Post) (*Post, error) {

	profile, err := s.GetProfile(post.Author.Id)
	if err != nil {
		log.Printf("Error retrieving user-profile, err : %s", err.Error())
	}
	_ = profile

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

func (s *server) GetPost(id int) (*Post, error) {
	panic("implement me")
}

func (s *server) ListPost() ([]*Post, error) {
	panic("implement me")
}

func (s *server) DeletePost(id int) error {
	panic("implement me")
}

func (s *server) SearchPost(params []string) ([]*Post, error) {
	panic("implement me")
}

func (s *server) CreateProfile(firstName, lastName, email , password string, opts ...string) (*Profile, error) {
	panic("implement me")
}

func (s *server) GetProfile(id string) (*Profile, error) {
	panic("implement me")
}

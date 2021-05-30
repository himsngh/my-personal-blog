package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strconv"
	"time"
)

type dbStore struct {
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

	return &dbStore{
		db:database,
	}, nil
}

func (s *dbStore) CreatePost(post *Post) (*Post, error) {

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

func (s *dbStore) GetPost(id int) (*Post, error) {
	panic("implement me")
}

func (s *dbStore) ListPost() ([]*Post, error) {
	//panic("implement me")

	posts := &Post{
		Id:     "123",
		Title:  "Test",
		Data:   "Testing",
		Date:   time.Time{}.Local(),
		Author: nil,
	}

	return []*Post{posts}, nil
}

func (s *dbStore) DeletePost(id int) error {
	panic("implement me")
}

func (s *dbStore) SearchPost(params []string) ([]*Post, error) {
	panic("implement me")
}

func (s *dbStore) CreateProfile(firstName, lastName, email , password string) (*Profile, error) {
	panic("implement me")
}

func (s *dbStore) GetProfile(id string) (*Profile, error) {
	panic("implement me")
}

func (s *dbStore) DeleteProfile(id string) {
	panic("implement me")
}

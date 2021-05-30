package database

import (
	"time"
)

type Store interface {
	CreatePost(post *Post) (*Post, error)
	GetPost(id int) (*Post, error)
	ListPost() ([]*Post, error)
	DeletePost(id int) error
	SearchPost(params []string) ([]*Post, error)

	CreateProfile(firstName, lastName, email , password string, opts ...string) (*Profile, error)
	GetProfile(id string) (*Profile, error)
	DeleteProfile(id string)

}

type Post struct {
	Id     string
	Title  string
	Data   string
	Date   time.Time
	Author *Profile
}

type Profile struct {
	Id             string
	FirstName      string
	LastName       string
	Email          string
	ProfilePicture string
}

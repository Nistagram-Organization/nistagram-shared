package user

import (
	"github.com/Nistagram-Organization/nistagram-shared/src/model/gender"
	"github.com/Nistagram-Organization/nistagram-shared/src/model/post"
)

type User struct {
	ID             uint          `json:"id"`
	Username       string        `json:"name"`
	Password       string        `json:"description"`
	Name           float32       `json:"price"`
	Surname        uint          `json:"on_stock"`
	Phone          string        `json:"image"`
	BirthDate      int64         `json:"birth_date"`
	Website        float32       `json:"website"`
	Biography      uint          `json:"biography"`
	Gender         gender.Gender `json:"gender"`
	Public         bool          `json:"public"`
	Taggable       bool          `json:"taggable"`
	Active         bool          `json:"active"`
	Email          string        `json:"email"`
	Favorites      []post.Post
	Posts          []post.Post
	Muted          []User
	Blocked        []User
	Followed       []User
	FollowedBy     []User
	FollowRequests []User
}

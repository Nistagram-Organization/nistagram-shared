package user

import (
	"github.com/Nistagram-Organization/nistagram-shared/src/model/gender"
	"github.com/Nistagram-Organization/nistagram-shared/src/model/post"
)

type User struct {
	ID             uint          `json:"id"`
	Username       string        `json:"username"`
	Password       string        `json:"password"`
	Name           string        `json:"name"`
	Surname        string        `json:"surname"`
	Phone          string        `json:"phone"`
	BirthDate      int64         `json:"birth_date"`
	Website        string        `json:"website"`
	Biography      string        `json:"biography"`
	Gender         gender.Gender `json:"gender"`
	Public         bool          `json:"public"`
	Taggable       bool          `json:"taggable"`
	Active         bool          `json:"active"`
	Email          string        `json:"email"`
	Favorites      []post.Post
	Posts          []post.Post
	Muted          []User        `gorm:"many2many:user_muted"`
	Blocked        []User        `gorm:"many2many:user_blocked"`
	Following      []User        `gorm:"many2many:user_following"`
	FollowedBy     []User        `gorm:"many2many:user_following"`
	FollowRequests []User        `gorm:"many2many:user_followRequests"`
}

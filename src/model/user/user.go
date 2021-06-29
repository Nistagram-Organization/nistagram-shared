package user

import (
	"github.com/Nistagram-Organization/nistagram-shared/src/model/gender"
	"github.com/Nistagram-Organization/nistagram-shared/src/model/post"
)

type User struct {
	ID             uint            `json:"id"`
	Username       string          `json:"username" gorm:"unique"`
	Name           string          `json:"name"`
	Surname        string          `json:"surname"`
	Phone          string          `json:"phone"`
	BirthDate      int64           `json:"birth_date"`
	Website        string          `json:"website"`
	Biography      string          `json:"biography"`
	Gender         gender.Gender   `json:"gender"`
	Public         bool            `json:"public"`
	Taggable       bool            `json:"taggable"`
	Active         bool            `json:"active"`
	Email          string          `json:"email" gorm:"unique"`
	Favorites      []post.PostUser `gorm:"many2many:favorites;joinForeignKey:UserID;JoinReferences:PostID"`
	Posts          []post.PostUser `gorm:"foreignKey:UserID"`
	Muted          []User          `gorm:"many2many:user_muted;joinForeignKey:UserID;JoinReferences:MutedUser"`
	Blocked        []User          `gorm:"many2many:user_blocked;joinForeignKey:UserID;JoinReferences:BlockedUser"`
	Following      []User          `gorm:"many2many:user_following;joinForeignKey:UserID;JoinReferences:FollowedUser"`
	FollowedBy     []User          `gorm:"many2many:user_following;joinForeignKey:UserID;JoinReferences:FollowedUser"`
	FollowRequests []User          `gorm:"many2many:user_followRequests;joinForeignKey:UserID;JoinReferences:FollowRequestsUser"`
}

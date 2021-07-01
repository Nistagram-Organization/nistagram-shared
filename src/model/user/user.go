package user

import (
	"github.com/Nistagram-Organization/nistagram-shared/src/model/gender"
	"github.com/Nistagram-Organization/nistagram-shared/src/model/post"
	"github.com/Nistagram-Organization/nistagram-shared/src/utils/rest_error"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
)

type User struct {
	ID             uint
	Name           string
	OwnerID        uint
	OwnerType      string
	Username       string          `json:"username" gorm:"unique"`
	FirstName      string          `json:"first_name"`
	LastName       string          `json:"last_name"`
	Phone          string          `json:"phone"`
	BirthDate      int64           `json:"birth_date"`
	Website        string          `json:"website"`
	Biography      string          `json:"biography"`
	Gender         gender.Gender   `json:"gender"`
	Public         bool            `json:"public"`
	Taggable       bool            `json:"taggable"`
	Active         bool            `json:"active"`
	Email          string          `json:"email" gorm:"unique"`
	Favorites      []post.PostUser `gorm:"many2many:favorites;joinForeignKey:UserID;JoinReferences:PostUserID"`
	Muted          []User          `gorm:"many2many:user_muted;joinForeignKey:UserID;JoinReferences:MutedUser"`
	Blocked        []User          `gorm:"many2many:user_blocked;joinForeignKey:UserID;JoinReferences:BlockedUser"`
	Following      []User          `gorm:"many2many:user_following;joinForeignKey:UserID;JoinReferences:FollowedUser"`
	FollowedBy     []User          `gorm:"many2many:user_following;joinForeignKey:UserID;JoinReferences:FollowedUser"`
	FollowRequests []User          `gorm:"many2many:user_followRequests;joinForeignKey:UserID;JoinReferences:FollowRequestsUser"`
}

func (u *User) Validate() rest_error.RestErr {
	if _, err := mail.ParseAddress(u.Email); err != nil {
		return rest_error.NewBadRequestError("Invalid email address")
	}
	if strings.TrimSpace(u.Username) == "" || len(u.Username) < 4 {
		return rest_error.NewBadRequestError("Username must be at least 4 characters long")
	}
	if strings.TrimSpace(u.FirstName) == "" {
		return rest_error.NewBadRequestError("First name cannot be empty")
	}
	if strings.TrimSpace(u.LastName) == "" {
		return rest_error.NewBadRequestError("Last name cannot be empty")
	}
	if match, _ := regexp.MatchString("^[+]*[(]{0,1}[0-9]{1,4}[)]{0,1}[-\\s./0-9]*$", u.Phone); !match {
		return rest_error.NewBadRequestError("Phone does not meet required pattern")
	}
	if time.Unix(u.BirthDate, 0).After(time.Now()) {
		return rest_error.NewBadRequestError("Birth date must be in the past")
	}
	if u.OwnerType == "agents" {
		if _, err := url.ParseRequestURI(u.Website); err != nil {
			return rest_error.NewBadRequestError("Invalid website url")
		}
	}
	return nil
}

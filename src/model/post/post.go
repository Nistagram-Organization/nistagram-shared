package post

import (
	"github.com/Nistagram-Organization/nistagram-shared/src/model/user"
)

type Post struct {
	ID                    uint   `json:"id"`
	Description           string `json:"description"`
	Date                  int64  `json:"date"`
	MarkedAsInappropriate bool   `json:"marked_as_inappropriate"`
	UserID                uint
	MediaID				  uint
	TaggedUsers           []user.User
}

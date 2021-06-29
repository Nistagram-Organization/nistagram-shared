package registered_user

import (
	"github.com/Nistagram-Organization/nistagram-shared/src/model/user"
)

type RegisteredUser struct {
	ID             uint            `json:"id"`
	User           user.User       `gorm:"polymorphic:Owner;"`
}

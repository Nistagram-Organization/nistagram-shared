package campaign

import (
	"github.com/Nistagram-Organization/nistagram-shared/src/model/commercial"
	"github.com/Nistagram-Organization/nistagram-shared/src/model/gender"
)

type Campaign struct {
	ID          uint                    `json:"id"`
	StartAge    uint                    `json:"start_age"`
	EndAge      uint                    `json:"end_age"`
	Gender      gender.Gender           `json:"gender"`
	Commercials []commercial.Commercial `gorm:"foreignKey:ID"`
}

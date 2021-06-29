package single_campaign

import (
	"github.com/Nistagram-Organization/nistagram-shared/src/model/campaign"
)

type SingleCampaign struct {
	ID                uint  `json:"id"`
	Date              int64 `json:"date"`
	campaign.Campaign `gorm:"embedded"`
}

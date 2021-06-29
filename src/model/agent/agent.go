package agent

import (
	"github.com/Nistagram-Organization/nistagram-shared/src/model/campaign"
	"github.com/Nistagram-Organization/nistagram-shared/src/model/user"
)

type Agent struct {
	Campaigns []campaign.Campaign
	User      user.User `gorm:"embedded"`
}

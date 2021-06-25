package agent

import (
	"github.com/Nistagram-Organization/nistagram-shared/src/model/campaign"
	"github.com/Nistagram-Organization/nistagram-shared/src/model/user"
)

type Agent struct {
	ID        uint                `json:"id"`
	Campaigns []campaign.Campaign
	user.User
}

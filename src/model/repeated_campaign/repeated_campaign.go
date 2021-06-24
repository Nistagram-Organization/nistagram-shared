package repeated_campaign

import (
	"github.com/Nistagram-Organization/nistagram-shared/src/model/campaign"
)

type RepeatedCampaign struct {
	ID        uint  `json:"id"`
	Start     int64 `json:"start"`
	End       int64 `json:"end"`
	Frequency uint `json:"frequency"`
	campaign.Campaign
}

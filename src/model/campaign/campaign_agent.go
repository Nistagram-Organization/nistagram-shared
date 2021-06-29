package campaign

type CampaignAgent struct {
	ID         uint `json:"id"`
	CampaignID uint `json:"campaign_id"`
	AgentID    uint
}

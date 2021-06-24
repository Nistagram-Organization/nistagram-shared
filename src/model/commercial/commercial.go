package commercial

type Commercial struct {
	ID         uint   `json:"id"`
	Link       string `json:"link"`
	MediaID    uint
	CampaignID uint
}

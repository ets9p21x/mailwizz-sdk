package model

type CampaignRecord struct {
	CampaignUid string `json:"campaign_uid"`
	Name        string `json:"name"`
	Status      string `json:"status"`
}

type CampaignData struct {
	Count       string          `json:"count"`
	TotalPages  int             `json:"total_pages"`
	CurrentPage int             `json:"current_page"`
	Records     []CampaignRecord `json:"records"`
}

type Campaign struct {
	Status string      `json:"status"`
	Data   CampaignData `json:"data"`
	Error  string      `json:"error"`
}

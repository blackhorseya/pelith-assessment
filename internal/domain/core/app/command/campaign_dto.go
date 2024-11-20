package command

import (
	"time"
)

// createCampaignCommand represents the input data to create a campaign.
type createCampaignCommand struct {
	Name      string    `json:"name"`
	StartTime time.Time `json:"start_time"`
}

// campaignResponse represents the response data of a campaign.
type campaignResponse struct {
	ID string `json:"id"`
}

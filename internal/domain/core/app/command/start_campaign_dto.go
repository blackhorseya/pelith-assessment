package command

// StartCampaignCommand is the command for starting a campaign.
type StartCampaignCommand struct {
	ID string `json:"id"`
}

func (cmd StartCampaignCommand) Key() int {
	return startCampaignCommandKey
}

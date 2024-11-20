package command

import (
	"errors"
	"time"
)

// createCampaignCommand represents the input data to create a campaign.
type createCampaignCommand struct {
	Name      string    `json:"name"`
	StartTime time.Time `json:"start_time"`
}

func (cmd createCampaignCommand) Key() int {
	return createCampaignCommandKey
}

// Validate is used to validate the createCampaignCommand.
func (cmd createCampaignCommand) Validate() error {
	if cmd.Name == "" {
		return errors.New("name is required")
	}

	if cmd.StartTime.IsZero() {
		return errors.New("start time is required")
	}

	return nil
}

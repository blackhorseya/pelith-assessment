package command

import (
	"errors"
	"time"
)

// CreateCampaignCommand represents the input data to create a campaign.
type CreateCampaignCommand struct {
	Name      string         `json:"name"`
	StartTime time.Time      `json:"start_time"`
	Tasks     []*TaskCommand `json:"tasks"`
}

func (cmd CreateCampaignCommand) Key() int {
	return createCampaignCommandKey
}

// Validate is used to validate the CreateCampaignCommand.
func (cmd CreateCampaignCommand) Validate() error {
	if cmd.Name == "" {
		return errors.New("name is required")
	}

	if cmd.StartTime.IsZero() {
		return errors.New("start time is required")
	}

	return nil
}

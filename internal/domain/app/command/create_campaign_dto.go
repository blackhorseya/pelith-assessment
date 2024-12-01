package command

import (
	"errors"
	"time"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
)

// CreateCampaignCommand represents the input data to create a campaign.
type CreateCampaignCommand struct {
	Name       string             `json:"name"`
	StartTime  time.Time          `json:"start_time"`
	Mode       model.CampaignMode `json:"mode"`
	TargetPool string             `json:"target_pool"`
	MinAmount  float64            `json:"min_amount"`
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

//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package command

import (
	"context"
	"errors"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
)

type (
	// CampaignCreator is used to create a new campaign.
	CampaignCreator interface {
		// Create is used to create a new campaign.
		Create(c context.Context, campaign *model.Campaign) error
	}

	// CampaignUpdater is used to update the campaign.
	CampaignUpdater interface {
		// Update is used to update the campaign.
		Update(c context.Context, campaign *model.Campaign) error
	}
)

// CreateCampaignHandler handles the creation of a new campaign.
type CreateCampaignHandler struct {
	service biz.CampaignService
	repo    CampaignCreator
}

// NewCreateCampaignHandler is used to create a new CreateCampaignHandler.
func NewCreateCampaignHandler(service biz.CampaignService, repo CampaignCreator) *CreateCampaignHandler {
	return &CreateCampaignHandler{service: service, repo: repo}
}

// Handle is used to handle the creation of a new campaign.
func (h *CreateCampaignHandler) Handle(c context.Context, cmd createCampaignCommand) (string, error) {
	// validate the command
	if cmd.Name == "" {
		return "", errors.New("campaign name cannot be empty")
	}
	if cmd.StartTime.IsZero() {
		return "", errors.New("campaign start time is required")
	}

	// call domain service to start a new campaign
	campaign, err := h.service.StartCampaign(c, cmd.Name, cmd.StartTime)
	if err != nil {
		return "", err
	}

	// save the campaign to the repository
	err = h.repo.Create(c, campaign)
	if err != nil {
		return "", err
	}

	return campaign.Id, nil
}

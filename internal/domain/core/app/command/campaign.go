//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package command

import (
	"context"
	"errors"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/internal/shared/usecase"
)

type (
	// CampaignCreator is used to create a new campaign.
	CampaignCreator interface {
		// Create is used to create a new campaign.
		Create(c context.Context, campaign *biz.Campaign) error
	}

	// CampaignUpdater is used to update the campaign.
	CampaignUpdater interface {
		// Update is used to update the campaign.
		Update(c context.Context, campaign *biz.Campaign) error
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

func (h *CreateCampaignHandler) Handle(c context.Context, msg usecase.Message) error {
	cmd, ok := msg.(CreateCampaignCommand)
	if !ok {
		return errors.New("cannot handle the message")
	}

	err := cmd.Validate()
	if err != nil {
		return err
	}

	campaign, err := h.service.StartCampaign(c, cmd.Name, cmd.StartTime, nil)
	if err != nil {
		return err
	}

	return h.repo.Create(c, campaign)
}

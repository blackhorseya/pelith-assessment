package command

import (
	"context"
	"errors"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/internal/domain/repo"
	"github.com/blackhorseya/pelith-assessment/internal/shared/usecase"
)

// CreateCampaignHandler handles the creation of a new campaign.
type CreateCampaignHandler struct {
	service         biz.CampaignService
	repo            CampaignCreator
	campaignCreator repo.CampaignCreator
}

// NewCreateCampaignHandler is used to create a new CreateCampaignHandler.
func NewCreateCampaignHandler(
	service biz.CampaignService,
	repo CampaignCreator,
	campaignCreator repo.CampaignCreator,
) *CreateCampaignHandler {
	return &CreateCampaignHandler{
		service:         service,
		repo:            repo,
		campaignCreator: campaignCreator,
	}
}

func (h *CreateCampaignHandler) Handle(c context.Context, msg usecase.Message) (string, error) {
	cmd, ok := msg.(CreateCampaignCommand)
	if !ok {
		return "", errors.New("cannot handle the message")
	}

	err := cmd.Validate()
	if err != nil {
		return "", err
	}

	campaign, err := h.service.CreateCampaign(c, cmd.Name, cmd.StartTime, cmd.Mode, cmd.TargetPool, cmd.MinAmount)
	if err != nil {
		return "", err
	}

	err = h.repo.Create(c, campaign)
	if err != nil {
		return "", err
	}

	return campaign.Id, nil
}

package repo

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/entity/domain/core/model"
)

// ListCampaignCondition is a struct to define the condition of list campaign
type ListCampaignCondition struct {
}

type (
	// CampaignCreator is an interface to create a new campaign
	CampaignCreator interface {
		Create(c context.Context, campaign *biz.Campaign) error
	}

	CampaignUpdater interface {
		UpdateStatus(c context.Context, id string, status model.CampaignStatus) error
	}

	CampaignFinder interface {
		GetByID(c context.Context, id string) (*biz.Campaign, error)
		List(c context.Context, cond ListCampaignCondition) ([]*biz.Campaign, int64, error)
	}
)

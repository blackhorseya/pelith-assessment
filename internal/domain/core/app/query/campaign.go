//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package query

import (
	"context"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
)

// ListCampaignCondition is the condition to list the campaign.
type ListCampaignCondition struct {
}

// CampaignGetter is the interface that provides the methods to get the campaign.
type CampaignGetter interface {
	// GetByID is used to get a campaign by id.
	GetByID(c context.Context, id string) (*biz.Campaign, error)

	// List is used to list the campaign.
	List(c context.Context, cond ListCampaignCondition) (items []*biz.Campaign, total int, err error)
}

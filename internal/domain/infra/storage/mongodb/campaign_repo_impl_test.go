package mongodb

import (
	"context"
	"testing"
	"time"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/internal/shared/mongodbx"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)

type suiteCampaignRepoImpl struct {
	suite.Suite

	container *mongodbx.Container
	rw        *mongo.Client

	repo *CampaignRepoImpl
}

func (s *suiteCampaignRepoImpl) SetupTest() {
	container, err := mongodbx.NewContainer(context.Background())
	s.Require().NoError(err)
	s.container = container

	rw, err := container.RW(context.Background())
	s.Require().NoError(err)
	s.rw = rw

	s.repo = NewCampaignRepoImpl(rw)
}

func (s *suiteCampaignRepoImpl) TearDownTest() {
	if s.rw != nil {
		_ = s.rw.Disconnect(context.Background())
	}

	if s.container != nil {
		_ = s.container.Terminate(context.Background())
	}
}

func TestAllCampaignRepoImpl(t *testing.T) {
	suite.Run(t, new(suiteCampaignRepoImpl))
}

func (s *suiteCampaignRepoImpl) TestCampaignRepoImpl_Create() {
	campaign, _ := biz.NewCampaign("test", time.Now(), "0xB4e16d0168e52d35CaCD2c6185b44281Ec28C9Dc")

	type args struct {
		c        context.Context
		campaign *biz.Campaign
		mock     func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "create success",
			args:    args{campaign: campaign},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.c = context.Background()
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.repo.Create(tt.args.c, tt.args.campaign); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

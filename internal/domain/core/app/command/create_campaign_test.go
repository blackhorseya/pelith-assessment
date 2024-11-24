package command

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/blackhorseya/pelith-assessment/entity/domain/core/biz"
	"github.com/blackhorseya/pelith-assessment/internal/shared/usecase"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type suiteCampaignCommandTester struct {
	suite.Suite

	ctrl    *gomock.Controller
	repo    *MockCampaignCreator
	service *biz.MockCampaignService
	handler usecase.Handler
}

func (s *suiteCampaignCommandTester) SetupTest() {
	s.ctrl = gomock.NewController(s.T())

	s.repo = NewMockCampaignCreator(s.ctrl)
	s.service = biz.NewMockCampaignService(s.ctrl)

	s.handler = NewCreateCampaignHandler(s.service, s.repo)
}

func (s *suiteCampaignCommandTester) TearDownTest() {
	if s.ctrl != nil {
		s.ctrl.Finish()
	}
}

func TestCampaignCommandAll(t *testing.T) {
	suite.Run(t, new(suiteCampaignCommandTester))
}

func (s *suiteCampaignCommandTester) TestCreateCampaignHandler_Handle() {
	type args struct {
		c    context.Context
		msg  usecase.Message
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "invalid command",
			args: args{msg: CreateCampaignCommand{
				Name:      "",
				StartTime: time.Time{},
			}},
			wantErr: true,
		},
		{
			name: "start campaign failed",
			args: args{msg: CreateCampaignCommand{
				Name:      "test",
				StartTime: time.Now(),
			}, mock: func() {
				s.service.EXPECT().CreateCampaign(
					gomock.Any(),
					"test",
					gomock.Any(),
					gomock.Any(),
					gomock.Any(),
					gomock.Any(),
				).
					Return(nil, errors.New("mock error")).
					Times(1)
			}},
			wantErr: true,
		},
		{
			name: "save campaign failed",
			args: args{msg: CreateCampaignCommand{
				Name:      "test",
				StartTime: time.Now(),
			}, mock: func() {
				s.service.EXPECT().CreateCampaign(
					gomock.Any(),
					"test",
					gomock.Any(),
					gomock.Any(),
					gomock.Any(),
					gomock.Any(),
				).
					Return(&biz.Campaign{}, nil).
					Times(1)
				s.repo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(errors.New("mock error")).Times(1)
			}},
			wantErr: true,
		},
		{
			name: "success",
			args: args{msg: CreateCampaignCommand{
				Name:      "test",
				StartTime: time.Now(),
			}, mock: func() {
				s.service.EXPECT().CreateCampaign(
					gomock.Any(),
					"test",
					gomock.Any(),
					gomock.Any(),
					gomock.Any(),
					gomock.Any(),
				).
					Return(&biz.Campaign{}, nil).
					Times(1)
				s.repo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.c = context.Background()
			if tt.args.mock != nil {
				tt.args.mock()
			}

			_, err := s.handler.Handle(tt.args.c, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

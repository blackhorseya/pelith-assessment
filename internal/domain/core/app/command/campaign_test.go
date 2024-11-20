package command

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type suiteCampaignCommandTester struct {
	suite.Suite
}

func TestCampaignCommandAll(t *testing.T) {
	suite.Run(t, new(suiteCampaignCommandTester))
}

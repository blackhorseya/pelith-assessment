package pg

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type suiteUserRepoTester struct {
	suite.Suite
}

func TestUserRepoAll(t *testing.T) {
	suite.Run(t, new(suiteUserRepoTester))
}

package biz

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type suiteUserServiceTester struct {
	suite.Suite
}

func TestUserServiceAll(t *testing.T) {
	suite.Run(t, new(suiteUserServiceTester))
}

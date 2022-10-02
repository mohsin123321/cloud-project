package test_utilities

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mohsin123321/cloud-project/utility"
)

func setupUtility(t *testing.T) utility.Utility {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ut := utility.Utility{}

	return ut
}

package test_dataservice

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mohsin123321/cloud-project/dataservice"
	"github.com/mohsin123321/cloud-project/tests/mock_interfaces"
)

type DataserviceMockedComp struct {
	Db *mock_interfaces.MockDatabaseInterface
}

// setup the mocked interfaces used in testing
func setup(t *testing.T) (dataservice.Dataservice, DataserviceMockedComp) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	db := mock_interfaces.NewMockDatabaseInterface(mockCtrl)

	ds := dataservice.Dataservice{
		Db: db,
	}

	mocks := DataserviceMockedComp{
		Db: db,
	}

	return ds, mocks
}

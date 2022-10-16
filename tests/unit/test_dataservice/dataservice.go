package test_dataservice

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mohsin123321/cloud-project/dataservice"
	"github.com/mohsin123321/cloud-project/model"
	"github.com/mohsin123321/cloud-project/tests/mock_interfaces"
	"github.com/mohsin123321/cloud-project/tests/unit/mock_data"
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

// Matching that object is copied properly as expected
// and id is properly initialized
func MatchDataBody(data model.Data) bool {
	expected := model.Data{
		ID:        data.ID,
		DeviceID:  mock_data.DataBody.DeviceID,
		Value:     mock_data.DataBody.Value,
		Type:      mock_data.DataBody.Type,
		Date:      mock_data.DataBody.Time,
		Latitude:  mock_data.DataBody.Latitude,
		Longitude: mock_data.DataBody.Longitude,
	}
	return reflect.DeepEqual(expected, data) && !data.ID.IsZero()
}

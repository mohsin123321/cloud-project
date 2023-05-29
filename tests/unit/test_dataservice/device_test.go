package test_dataservice

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mohsin123321/cloud-project/model"
	"github.com/mohsin123321/cloud-project/tests/unit/mock_data"
	"gopkg.in/go-playground/assert.v1"
)

func TestInsertData_Success(t *testing.T) {
	ds, mocks := setup(t)
	expected := mock_data.DeviceDatBody()

	mocks.Db.EXPECT().InsertData(gomock.AssignableToTypeOf(model.Data{})).DoAndReturn(func(actual model.Data) error {
		// ID is properly assigned with value
		assert.Equal(t, true, !actual.ID.IsZero())
		// compare the rest of the fields with the actual value
		assert.Equal(t, expected.DeviceID, actual.DeviceID)
		assert.Equal(t, expected.Value, actual.Value)
		assert.Equal(t, expected.Type, actual.Type)
		assert.Equal(t, expected.Time, actual.Date)
		assert.Equal(t, expected.Latitude, actual.Latitude)
		assert.Equal(t, expected.Longitude, actual.Longitude)
		return nil
	})

	err := ds.InsertData(expected)
	assert.Equal(t, err, nil)
}

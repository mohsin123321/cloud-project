package test_dataservice

import (
	"testing"

	"github.com/mohsin123321/cloud-project/tests/unit/mock_data"
	"github.com/stretchr/testify/mock"
)

func TestInsertData_Success(t *testing.T) {
	ds, mocks := setup(t)
	mocks.Db.EXPECT().InsertData(mock.MatchedBy(MatchDataBody))

	ds.InsertData(mock_data.DataBody)
}

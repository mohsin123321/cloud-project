package test_controller

import (
	"net/http"
	"testing"

	"github.com/mohsin123321/cloud-project/dto"
	"github.com/mohsin123321/cloud-project/tests/unit/mock_data"
)

func TestInsertData_Success(t *testing.T) {
	// setup mock interfaces
	ctrl, mocks := setup(t)

	// making the request
	body := mock_data.DataBody
	path := "/device/123"

	req := createRequest(t, http.MethodGet, path, body)

	// setup expected function calls
	mocks.Ut.EXPECT().ParseBody(req.Body, &dto.PostDataBody{}).Do(mockParseBody)
	mocks.Service.EXPECT().InsertData(body)
	resp := serveRequest(ctrl.InsertData, req)

	checkResponse(t, resp, http.StatusNoContent, nil, nil)
}

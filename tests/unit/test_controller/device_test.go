package test_controller

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/mohsin123321/cloud-project/dto"
	"github.com/mohsin123321/cloud-project/error_handling"
	"github.com/mohsin123321/cloud-project/tests/unit/mock_data"
)

func TestInsertData_Success(t *testing.T) {
	// setup mock interfaces
	ctrl, mocks, w := setup(t)

	// making the request
	body := mock_data.DeviceDatBody()
	path := "/device"

	req := createRequest(t, http.MethodPost, path, body)

	// setup expected function calls
	mocks.Ut.EXPECT().ParseBody(req.Body, &dto.PostDataBody{}).Do(func(body io.ReadCloser, dest interface{}) {
		_ = json.NewDecoder(body).Decode(&dest)
	})
	mocks.Service.EXPECT().InsertData(body).Return(nil)
	mocks.Ut.EXPECT().EncodeEmptyResponse(req, w, nil)

	ctrl.InsertData(w, req)
}

func TestInsertData_BadSyntax(t *testing.T) {
	// setup mock interfaces
	ctrl, mocks, w := setup(t)

	// making the request
	body := mock_data.DeviceDatBody()
	path := "/device"

	req := createRequest(t, http.MethodPost, path, body)
	err := error_handling.ErrBadSyntax()

	// setup expected function calls
	mocks.Ut.EXPECT().ParseBody(req.Body, &dto.PostDataBody{}).Return(error_handling.ErrBadSyntax())
	mocks.Ut.EXPECT().EncodeErrResponse(req, w, err)

	ctrl.InsertData(w, req)
}

package test_controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mohsin123321/cloud-project/controller"
	"github.com/mohsin123321/cloud-project/tests/mock_interfaces"
)

type ControllerMockedComp struct {
	Service *mock_interfaces.MockDataserviceInterface
	Ut      *mock_interfaces.MockUtilityInterface
}

func setup(t *testing.T) (controller.HttpController, ControllerMockedComp, http.ResponseWriter) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ds := mock_interfaces.NewMockDataserviceInterface(mockCtrl)
	ut := mock_interfaces.NewMockUtilityInterface(mockCtrl)

	ctrl := controller.HttpController{
		Ds: ds,
		Ut: ut,
	}

	mocks := ControllerMockedComp{
		Service: ds,
		Ut:      ut,
	}

	return ctrl, mocks, httptest.NewRecorder()
}

// // serveRequest create a newRecorder and pass it to the f
// func serveRequest(f func(http.ResponseWriter, *http.Request), req *http.Request) *http.Response {

// 	w := httptest.NewRecorder()
// 	f(w, req)

// 	return w.Result()
// }

// createRequest generates the http request with the method, path, body and vars(params passed into the url)
func createRequest(t *testing.T, method string, path string, body interface{}) *http.Request {
	var req *http.Request

	if body != nil {
		requestBody, err := json.Marshal(body)
		if err != nil {
			t.Fatalf("Error marshaling request body: %v", err)
		}

		req = httptest.NewRequest(method, path, bytes.NewReader(requestBody))
	} else {
		req = httptest.NewRequest(method, path, nil)
	}

	return req
}

// // checkResponse check the statusCode and the body of the response
// func checkResponse(t *testing.T, resp *http.Response, expectedHttpStatus int, expectedBody interface{}, bodyTypeExp interface{}) {

// 	defer resp.Body.Close()

// 	// check the statusCode
// 	assert.Equal(t, expectedHttpStatus, resp.StatusCode)

// 	// check the body
// 	if expectedBody != nil {

// 		b, err := io.ReadAll(resp.Body)
// 		if err != nil {
// 			t.Error()
// 		}

// 		err = json.Unmarshal(b, &bodyTypeExp)
// 		if err != nil {
// 			t.Error()
// 		}

// 		assert.Equal(t, expectedBody, bodyTypeExp)
// 	}

// }

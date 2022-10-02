package test_controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/mohsin123321/cloud-project/controller"
	"github.com/mohsin123321/cloud-project/error_handling"
	"github.com/mohsin123321/cloud-project/tests/mock_interfaces"
	"gopkg.in/go-playground/assert.v1"
)

type ControllerMockedComp struct {
	Service *mock_interfaces.MockDataserviceInterface
	Ut      *mock_interfaces.MockUtilityInterface
}

func setup(t *testing.T) (controller.HttpController, ControllerMockedComp) {
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

	return ctrl, mocks
}

// serveRequest create a newRecorder and pass it to the f
func serveRequest(f func(http.ResponseWriter, *http.Request), req *http.Request) *http.Response {

	w := httptest.NewRecorder()
	f(w, req)

	return w.Result()
}

// createRequest generates the http request with the method, path, body and vars(params passed into the url)
func createRequest(method string, path string, body interface{}, vars map[string]string) *http.Request {
	var req *http.Request

	if body != nil {
		var requestBody []byte

		switch body.(type) {
		case string:
			requestBody = []byte(fmt.Sprint(body))
		default:
			if bytes, err := json.Marshal(body); err != nil {
				log.Fatal(err)
			} else {
				requestBody = bytes
			}
		}

		req = httptest.NewRequest(method, path, bytes.NewReader(requestBody))
	} else {
		req = httptest.NewRequest(method, path, nil)
	}

	return mux.SetURLVars(req, vars)
}

// checkResponse check the statusCode and the body of the response
func checkResponse(t *testing.T, resp *http.Response, expectedHttpStatus int, expectedBody interface{}, bodyTypeExp interface{}) {

	defer resp.Body.Close()

	// check the statusCode
	assert.Equal(t, expectedHttpStatus, resp.StatusCode)

	// check the body
	if expectedBody != nil {

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Error()
		}

		err = json.Unmarshal(b, &bodyTypeExp)
		if err != nil {
			t.Error()
		}

		assert.Equal(t, expectedBody, bodyTypeExp)
	}

}

// mockParseBody mocks the parseBody function
func mockParseBody(body io.ReadCloser, dest interface{}) bool {
	if body == http.NoBody {
		error_handling.ThrowError(error_handling.ErrBadSyntax)
	}

	// decode
	err := json.NewDecoder(body).Decode(&dest)
	if err != nil {
		error_handling.ThrowError(error_handling.ErrBadSyntax)
	}

	// validate
	v := validator.New()

	err = v.Struct(dest)
	if err != nil {
		error_handling.ThrowError(error_handling.ErrBadSyntax)
	}
	return true
}

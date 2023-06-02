package test_utilities

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mohsin123321/cloud-project/dto"
	"github.com/mohsin123321/cloud-project/error_handling"
	"github.com/mohsin123321/cloud-project/model"
	"github.com/mohsin123321/cloud-project/tests/unit/mock_data"
	"gopkg.in/go-playground/assert.v1"
)

func TestEncodeEmptyResponse(t *testing.T) {
	ut := setupUtility(t)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		panic(err)
	}

	rr := httptest.NewRecorder()

	ut.EncodeEmptyResponse(req, rr, nil)

	// check the response code
	assert.Equal(t, http.StatusNoContent, rr.Code)
}

func TestEncodeErrResponse(t *testing.T) {
	ut := setupUtility(t)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		panic(err)
	}

	ctx := context.WithValue(req.Context(), model.GetCtxKeyID(), mock_data.NewID())
	req = req.WithContext(ctx)

	e := error_handling.ErrServerError()
	rr := httptest.NewRecorder()
	errObj := dto.Error{
		Err: e.Error(),
		Msg: e.Msg(),
	}
	expected := mock_data.ToJSON(errObj)

	ut.EncodeErrResponse(req, rr, err)

	// check the expected response
	assert.Equal(t, expected, rr.Body.String())
	// check the response code
	assert.Equal(t, e.Status(), rr.Code)
}

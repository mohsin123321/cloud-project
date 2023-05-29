package test_utilities

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	"github.com/mohsin123321/cloud-project/error_handling"
	"gopkg.in/go-playground/assert.v1"
)

func TestParseBody_Fail_NoBody(t *testing.T) {
	ut := setupUtility(t)

	err := ut.ParseBody(io.NopCloser(bytes.NewReader(nil)), nil)

	assert.Equal(t, err, error_handling.ErrBadSyntax())
}

func TestParseBody_Fail_Decode_TypeErr(t *testing.T) {
	ut := setupUtility(t)

	type BodyErr struct {
		Value int
	}
	bodyErr := BodyErr{
		Value: 123,
	}
	jsonBody, _ := json.Marshal(bodyErr)

	type Body struct {
		Value string
	}
	actual := Body{}
	err := ut.ParseBody(io.NopCloser(bytes.NewReader(jsonBody)), &actual)

	assert.Equal(t, err, error_handling.ErrBadSyntax())
}

func TestParseBody_Fail_Required(t *testing.T) {
	ut := setupUtility(t)
	type Body struct {
		Value string `validate:"required"`
	}
	expected := Body{
		Value: "",
	}
	jsonBody, _ := json.Marshal(expected)

	actual := Body{}
	err := ut.ParseBody(io.NopCloser(bytes.NewReader(jsonBody)), &actual)

	assert.Equal(t, err, error_handling.ErrBadSyntax())
}

func TestParseBody_Success_Required(t *testing.T) {
	ut := setupUtility(t)

	type Body struct {
		Value string `validate:"required"`
	}

	expected := Body{
		Value: "test",
	}

	jsonBody, _ := json.Marshal(expected)

	actual := Body{}
	err := ut.ParseBody(io.NopCloser(bytes.NewReader(jsonBody)), &actual)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

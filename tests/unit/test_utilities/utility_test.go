package test_utilities

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	"github.com/mohsin123321/cloud-project/error_handling"
	"github.com/mohsin123321/cloud-project/tests/unit"
	"gopkg.in/go-playground/assert.v1"
)

func TestParseBody_Fail_NoBody(t *testing.T) {
	defer unit.RecoverError(error_handling.ErrBadSyntax, t)
	ut := setupUtility(t)

	ut.ParseBody(io.NopCloser(bytes.NewReader(nil)), nil)
	t.Error("No Panic found")
}

func TestParseBody_Fail_Decode_TypeErr(t *testing.T) {

	defer unit.RecoverError(error_handling.ErrBadSyntax, t)

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
	ut.ParseBody(io.NopCloser(bytes.NewReader(jsonBody)), &actual)
	t.Error("No Panic found")
}

func TestParseBody_Fail_Required(t *testing.T) {

	defer unit.RecoverError(error_handling.ErrBadSyntax, t)

	ut := setupUtility(t)
	type Body struct {
		Value string `validate:"required"`
	}
	expected := Body{
		Value: "",
	}
	jsonBody, _ := json.Marshal(expected)

	actual := Body{}
	ut.ParseBody(io.NopCloser(bytes.NewReader(jsonBody)), &actual)
	t.Error("No Panic found")
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
	ut.ParseBody(io.NopCloser(bytes.NewReader(jsonBody)), &actual)
	assert.Equal(t, expected, actual)
}

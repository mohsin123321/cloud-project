package utility

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/mohsin123321/cloud-project/error_handling"
)

type ParserInterface interface {
	// ParseBody decode and validate the body request and fills the provided interface with data
	ParseBody(body io.ReadCloser, dest interface{}) error
}

func (ut *Utility) ParseBody(body io.ReadCloser, dest interface{}) error {
	if body == http.NoBody {
		return error_handling.ErrBadSyntax()
	}

	// decode
	err := json.NewDecoder(body).Decode(&dest)
	if err != nil {
		return error_handling.ErrBadSyntax()
	}

	// validate
	v := validator.New()
	err = v.Struct(dest)
	if err != nil {
		return error_handling.ErrBadSyntax()
	}

	return nil
}

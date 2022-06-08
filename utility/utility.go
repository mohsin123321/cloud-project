package utility

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-playground/validator"
	e "github.com/mohsin123321/cloud-project/error_handling"
)

type UtilityInterface interface {
	// ParseBody decode and validate the body request and fills the provided interface with data
	ParseBody(body io.ReadCloser, dest interface{})
}

type Utility struct{}

func (ut *Utility) ParseBody(body io.ReadCloser, dest interface{}) {

	if body == http.NoBody {
		e.ThrowError(e.ErrBadSyntax)
	}

	// decode
	err := json.NewDecoder(body).Decode(&dest)
	if err != nil {
		e.ThrowError(e.ErrBadSyntax)
	}

	// validate
	v := validator.New()
	err = v.Struct(dest)
	if err != nil {
		e.ThrowError(e.ErrBadSyntax)
	}
}

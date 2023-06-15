package utility

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/mohsin123321/cloud-project/error_handling"
	"github.com/mohsin123321/cloud-project/model"
)

const (
	MimeTypeJSON = "application/json"
)

type EncoderInterface interface {
	// EncodeEmptyResponse with status 204 No Content and empty body if err is nil, otherwise an error response is sent.
	EncodeEmptyResponse(r *http.Request, w http.ResponseWriter, err error)

	// EncodeErrResponse writes error response to the response writer.
	EncodeErrResponse(r *http.Request, w http.ResponseWriter, err error)
}

func (ut *Utility) EncodeEmptyResponse(r *http.Request, w http.ResponseWriter, err error) {
	if err != nil {
		ut.EncodeErrResponse(r, w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (ut *Utility) EncodeErrResponse(r *http.Request, w http.ResponseWriter, err error) {
	err = error_handling.PropagateError(err, 2)

	id, ok := r.Context().Value(model.GetCtxKeyID()).(uuid.UUID)
	if !ok {
		panic(error_handling.ErrServerError)
	}

	appErr, ok := err.(*error_handling.CustomError)

	if !ok { // should never happen
		appErr = error_handling.ErrServerError()
	}

	appErr.ID = id
	appErr.Log(error_handling.LogMessageErrorResponse)

	json := appErr.MarhsalJSON()
	w.Header().Set("Content-Type", MimeTypeJSON)
	w.WriteHeader(appErr.Status())
	_, _ = w.Write(json)
}

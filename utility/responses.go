package utility

import (
	"encoding/json"
	"net/http"

	"github.com/mohsin123321/cloud-project/model"
)

// JSONokResponse make a ok response with a json body
func JSONokResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", model.MimeTypeJSON)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(data)
}

// NoContentResponse 204 makes an ok response without a json body
func NoContentResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

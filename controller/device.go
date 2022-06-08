package controller

import (
	"net/http"
	"time"

	"github.com/mohsin123321/cloud-project/model"
)

type PostDataBody struct {
	Device string  `json:"deviceId" validate:"required"`
	Value  float32 `json:"value" validate:"required"`
	Time   int64   `json:"time"`
}

func (ctrl *HttpController) InsertData(w http.ResponseWriter, r *http.Request) {
	var body PostDataBody
	ctrl.Ut.ParseBody(r.Body, &body)
	data := model.Data{
		Type:  "temperature",
		Value: body.Value,
		Date:  time.Now(),
	}
	ctrl.DS.InsertData(data)
}

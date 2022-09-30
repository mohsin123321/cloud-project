package controller

import (
	"net/http"

	"github.com/mohsin123321/cloud-project/dto"
	"github.com/mohsin123321/cloud-project/utility"
)

func (ctrl *HttpController) InsertData(w http.ResponseWriter, r *http.Request) {
	var body dto.PostDataBody
	ctrl.Ut.ParseBody(r.Body, &body)
	ctrl.Ds.InsertData(body)

	utility.NoContentResponse(w)
}

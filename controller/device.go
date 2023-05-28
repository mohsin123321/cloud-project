package controller

import (
	"net/http"

	"github.com/mohsin123321/cloud-project/dto"
)

// @Summary create sensors data.
// @ID InsertData
// @tags device
// @Accept json
// @Param X-Auth-Token header string true "jwt token"
// @Param Body body dto.PostDataBody true "contains the information related to device"
// @Success 204
// @Failure 400 "ERR_BAD_SYNTAX"
// @Failure 500 "ERR_INTERNAL_SERVER_ERROR"
// @Router /api/device [post]
func (ctrl *HttpController) InsertData(w http.ResponseWriter, r *http.Request) {
	var body dto.PostDataBody
	err := ctrl.Ut.ParseBody(r.Body, &body)
	if err != nil {
		ctrl.Ut.EncodeErrResponse(r, w, err)
		return
	}
	err = ctrl.Ds.InsertData(body)
	ctrl.Ut.EncodeEmptyResponse(r, w, err)
}

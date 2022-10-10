package controller

import (
	"net/http"

	"github.com/mohsin123321/cloud-project/dto"
	"github.com/mohsin123321/cloud-project/utility"
)

// CreateClassroom insert the coming data from simulator.
// @Summary create sensors data.
// @tags device
// @Accept json
// @Param Body body dto.PostDataBody true "contains the information related to device"
// @Success 200
// @Router /device [post]
func (ctrl *HttpController) InsertData(w http.ResponseWriter, r *http.Request) {
	var body dto.PostDataBody
	ctrl.Ut.ParseBody(r.Body, &body)
	ctrl.Ds.InsertData(body)

	utility.NoContentResponse(w)
}

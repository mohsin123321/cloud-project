package controller

import (
	"net/http"

	"github.com/mohsin123321/cloud-project/dataservice"
	"github.com/mohsin123321/cloud-project/utility"
)

type ControllerInterface interface {
	InsertData(w http.ResponseWriter, r *http.Request)
}

type HttpController struct {
	Ds dataservice.DataserviceInterface
	Ut utility.UtilityInterface
}

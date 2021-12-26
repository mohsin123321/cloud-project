package controller

import (
	"net/http"

	"github.com/mohsin123321/cloud-project/dataservice"
)

type ControllerInterface interface {
	CreateTemperature(w http.ResponseWriter, r *http.Request)
}

type HttpController struct {
	DS dataservice.DataserviceInterface
}

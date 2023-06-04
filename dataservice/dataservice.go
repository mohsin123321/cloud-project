package dataservice

import (
	"github.com/mohsin123321/cloud-project/database"
	"github.com/mohsin123321/cloud-project/error_handling"
)

// package that contains all the business logic
type DataserviceInterface interface {
	DeviceInterface
}

type Dataservice struct {
	Db database.DatabaseInterface
}

// SetupDGS (DataGatewayService) initializes the dgs and return it
func SetupDGS() *Dataservice {
	var dgs Dataservice

	dgs.Db = database.SetupDB()

	return &dgs
}

func handleError(err error) error {
	if err != nil {
		err = error_handling.PropagateError(err, 2)
	}
	return err
}

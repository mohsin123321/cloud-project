package dataservice

import "github.com/mohsin123321/cloud-project/database"

// package that contains all the business logic

type DataserviceInterface interface {
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

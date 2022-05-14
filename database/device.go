package database

import (
	e "github.com/mohsin123321/cloud-project/error_handling"
	"github.com/mohsin123321/cloud-project/model"
)

func (db *Database) InsertData(data model.Data) {
	_, error := db.DB.Database("iot").Collection("iot_data").InsertOne(
		db.DBContext,
		data,
	)
	e.CheckDbError(error)
}

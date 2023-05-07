package database

import (
	"context"

	e "github.com/mohsin123321/cloud-project/error_handling"
	"github.com/mohsin123321/cloud-project/model"
)

func (db *Database) InsertData(data model.Data) {
	_, error := db.DB.Collection("iot_data").InsertOne(
		context.TODO(),
		data,
	)
	e.CheckDbError(error)
}

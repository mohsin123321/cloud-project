package database

import (
	"context"

	"github.com/mohsin123321/cloud-project/model"
)

type DeviceInterface interface {
	InsertData(model.Data) error
}

func (db *Database) InsertData(data model.Data) error {
	_, err := db.DB.Collection("iot_data").InsertOne(
		context.TODO(),
		data,
	)
	return handleError(err)
}

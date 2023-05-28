package dataservice

import (
	"github.com/mohsin123321/cloud-project/dto"
	"github.com/mohsin123321/cloud-project/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeviceInterface interface {
	InsertData(dto.PostDataBody) error
}

func (ds *Dataservice) InsertData(body dto.PostDataBody) error {
	data := model.Data{
		ID:        primitive.NewObjectID(),
		DeviceID:  body.DeviceID,
		Value:     body.Value,
		Type:      body.Type,
		Date:      body.Time,
		Latitude:  body.Latitude,
		Longitude: body.Longitude,
	}
	err := ds.Db.InsertData(data)
	return handleError(err)
}

package dataservice

import (
	"github.com/mohsin123321/cloud-project/dto"
	"github.com/mohsin123321/cloud-project/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ds *Dataservice) InsertData(body dto.PostDataBody) {
	data := model.Data{
		ID:        primitive.NewObjectID(),
		DeviceID:  body.DeviceID,
		Value:     body.Value,
		Type:      body.Type,
		Date:      body.Time,
		Latitude:  body.Latitude,
		Longitude: body.Longitude,
	}
	ds.Db.InsertData(data)
}

package dataservice

import (
	"github.com/mohsin123321/cloud-project/dto"
	"github.com/mohsin123321/cloud-project/model"
)

func (ds *Dataservice) InsertData(body dto.PostDataBody) {
	data := model.Data{
		DeviceID: body.DeviceID,
		Value:    body.Value,
		Type:     body.Type,
		Date:     body.Time,
	}
	ds.Db.InsertData(data)
}

package dataservice

import "github.com/mohsin123321/cloud-project/model"

func (ds *Dataservice) InsertData(data model.Data) {
	ds.Db.InsertData(data)
}

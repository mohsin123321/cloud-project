package model

import (
	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Data struct {
	ID       primitive.ObjectID
	DeviceID uuid.UUID
	Value    float32
	Type     string
	Date     int64
}

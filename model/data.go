package model

import (
	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Data struct {
	ID        primitive.ObjectID `bson:"_id"`
	DeviceID  uuid.UUID          `bson:"deviceId"`
	Value     float32            `bson:"value"`
	Type      string             `bson:"type"`
	Date      int64              `bson:"date"`
	Latitude  float64            `bson:"latitude"`
	Longitude float64            `bson:"longitude"`
}

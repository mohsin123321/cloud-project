package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Data struct {
	ID    primitive.ObjectID
	Value float32
	Type  string
	Date  time.Time
}

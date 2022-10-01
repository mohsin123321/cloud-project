package dto

import "github.com/gofrs/uuid"

type PostDataBody struct {
	DeviceID uuid.UUID `json:"deviceId" validate:"required"`
	Type     string    `json:"type" validate:"required,max=20"`
	Value    float32   `json:"value" validate:"required,numeric"`
	Time     int64     `json:"time" validate:"required,numeric"`
}

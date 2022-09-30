package dto

import "github.com/gofrs/uuid"

type PostDataBody struct {
	DeviceID uuid.UUID `json:"deviceId" validate:"required"`
	Type     string    `json:"type" vaidate:"required"`
	Value    float32   `json:"value" validate:"required"`
	Time     int64     `json:"time"`
}

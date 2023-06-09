package dto

import "github.com/gofrs/uuid"

type PostDataBody struct {
	DeviceID  uuid.UUID `json:"deviceId" validate:"required" example:"cbffd6a4-c0be-4dd0-b8f4-ebd53433e7cd"`
	Type      string    `json:"type" validate:"required,max=20" example:"temperature"`
	Value     float64   `json:"value" validate:"required,numeric" example:"25.6"`
	Time      int64     `json:"time" validate:"required,numeric" example:"1641567600"`
	Latitude  float64   `json:"lat" validate:"required,numeric" example:"45.464203"`
	Longitude float64   `json:"lng" validate:"required,numeric" example:"9.189982"`
}

package mock_data

import (
	"time"

	"github.com/mohsin123321/cloud-project/dto"
)

var DataBody = dto.PostDataBody{
	DeviceID: ID(),
	Type:     "temperature",
	Value:    49.3,
	Time:     time.Now().Unix(),
}
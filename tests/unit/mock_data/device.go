package mock_data

import (
	"time"

	"github.com/mohsin123321/cloud-project/dto"
)

func DeviceDatBody() dto.PostDataBody {
	return dto.PostDataBody{
		DeviceID:  NewID(),
		Type:      "temperature",
		Value:     49.3,
		Time:      time.Now().Unix(),
		Latitude:  45.464203,
		Longitude: 9.189982,
	}
}

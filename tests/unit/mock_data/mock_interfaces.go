package mock_data

import "github.com/gofrs/uuid"

func ID() uuid.UUID {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return id
}

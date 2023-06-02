package mock_data

import (
	"encoding/json"

	"github.com/gofrs/uuid"
)

// NewID returns the randomly generated uuid
func NewID() uuid.UUID {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return id
}

// ToJSON converts the given value to a json string.
// It panics if an error occurs.
func ToJSON(v interface{}) string {
	s, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return string(s)
}

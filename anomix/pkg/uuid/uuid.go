package uuid

import (
	"github.com/google/uuid"
)

// GetUUID generate and return unique id
// the id will be always 32 byte long unique string
func GetUUID() string {
	return uuid.New().String()
}

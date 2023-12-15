package internal

import (
	"github.com/google/uuid"
)

func GetUuid() string {
	return uuid.NewString()
}

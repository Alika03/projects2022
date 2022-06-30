package utils

import uuid "github.com/satori/go.uuid"

func GenerateUuid() uuid.UUID {
	return uuid.Must(uuid.NewV4(), nil)
}

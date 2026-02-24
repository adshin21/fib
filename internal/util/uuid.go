// Package util provides generic utility
package util

import (
	"fmt"

	"github.com/google/uuid"
)

func GetUUID() (uuid.UUID, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to generate UUID v7: %w", err)
	}
	return id, nil
}

func GetUUIDString() (string, error) {
	id, err := GetUUID()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

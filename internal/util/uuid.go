// Package util provides generic utility
package util

import "github.com/google/uuid"

func GetUUID() (uuid.UUID, error) {
	return uuid.NewV7()
}

func GetUUIDString() (string, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

package util

import "github.com/google/uuid"

func NewRandomUUID() (string, error) {
	randUUID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return randUUID.String(), nil
}

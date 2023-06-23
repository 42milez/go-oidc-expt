package xutil

import (
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/argon2"
)

const (
	errFailedToGenerateRandomBytes Err = "failed to generate random bytes"
)

type Err string

func (v Err) Error() string {
	return string(v)
}

func GeneratePasswordHash(pw string) (string, error) {
	// The parameters below are recommended in https://datatracker.ietf.org/doc/html/draft-irtf-cfrg-argon2-13 as SECOND
	// RECOMMENDED option.
	var memory uint32 = 64 * 1024
	var iterations uint32 = 3
	var parallelism uint8 = 4
	var saltLength uint32 = 128
	var keyLength uint32 = 256

	salt := make([]byte, saltLength)

	_, err := rand.Read(salt)

	if err != nil {
		return "", fmt.Errorf("%w, %w", errFailedToGenerateRandomBytes, err)
	}

	hash := argon2.IDKey([]byte(pw), salt, iterations, memory, parallelism, keyLength)

	return string(hash), nil
}

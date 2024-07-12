package security

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/gob"

	"golang.org/x/crypto/argon2"
)

type argon2Variant string

const (
	a2   argon2Variant = "argon2"
	a2i  argon2Variant = "argon2i"
	a2id argon2Variant = "argon2id"
)

type argon2Representation struct {
	Variant     argon2Variant
	Version     int
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	KeyLength   uint32
	Salt        []byte
	Hash        []byte
}

// The parameters below are recommended in https://datatracker.ietf.org/doc/html/draft-irtf-cfrg-argon2-13 as SECOND
// RECOMMENDED option.
const memory uint32 = 64 * 1024
const iterations uint32 = 3
const parallelism uint8 = 4
const saltLength uint32 = 16
const keyLength uint32 = 32

func HashPassword(raw string) (string, error) {
	salt := make([]byte, saltLength)

	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(raw), salt, iterations, memory, parallelism, keyLength)

	rep := argon2Representation{
		Variant:     a2id,
		Version:     argon2.Version,
		Memory:      memory,
		Iterations:  iterations,
		Parallelism: parallelism,
		KeyLength:   keyLength,
		Salt:        salt,
		Hash:        hash,
	}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	if err := enc.Encode(rep); err != nil {
		return "", err
	}

	ret := base64.RawStdEncoding.EncodeToString(buf.Bytes())

	return ret, nil
}

func ComparePassword(raw string, encoded string) (bool, error) {
	b, err := base64.RawStdEncoding.DecodeString(encoded)

	if err != nil {
		return false, err
	}

	buf := bytes.NewBuffer(b)
	dec := gob.NewDecoder(buf)
	repr := &argon2Representation{}

	if err = dec.Decode(&repr); err != nil {
		return false, err
	}

	hash := argon2.IDKey([]byte(raw), repr.Salt, repr.Iterations, repr.Memory, repr.Parallelism, repr.KeyLength)

	return bytes.Equal(hash, repr.Hash), nil
}

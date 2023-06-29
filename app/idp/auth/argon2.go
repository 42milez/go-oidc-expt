package auth

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/gob"

	"github.com/42milez/go-oidc-server/pkg/xerr"

	"github.com/42milez/go-oidc-server/app/idp/ent/typedef"
	"golang.org/x/crypto/argon2"
)

type argon2Variant string

const (
	a2   argon2Variant = "argon2"
	a2i  argon2Variant = "argon2i"
	a2id argon2Variant = "argon2id"
)

type passwordHashRepr struct {
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
const saltLength uint32 = 128
const keyLength uint32 = 256

func GeneratePasswordHash(raw string) (typedef.PasswordHash, error) {
	salt := make([]byte, saltLength)

	if _, err := rand.Read(salt); err != nil {
		return "", xerr.WrapErr(xerr.FailedToGenerateRandomBytes, err)
	}

	hash := argon2.IDKey([]byte(raw), salt, iterations, memory, parallelism, keyLength)

	repr := passwordHashRepr{
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

	if err := enc.Encode(repr); err != nil {
		return "", xerr.WrapErr(xerr.FailedToEncodeInToBytes, err)
	}

	ret := base64.RawStdEncoding.EncodeToString(buf.Bytes())

	return typedef.PasswordHash(ret), nil
}

func ComparePassword(raw string, encoded typedef.PasswordHash) (bool, error) {
	b, err := base64.RawStdEncoding.DecodeString(string(encoded))

	if err != nil {
		return false, xerr.WrapErr(xerr.FailedToDecodeInToBytes, err)
	}

	buf := bytes.NewBuffer(b)
	dec := gob.NewDecoder(buf)
	repr := &passwordHashRepr{}

	if err = dec.Decode(&repr); err != nil {
		return false, xerr.WrapErr(xerr.FailedToDecodeInToStruct, err)
	}

	hash := argon2.IDKey([]byte(raw), repr.Salt, repr.Iterations, repr.Memory, repr.Parallelism, repr.KeyLength)

	return bytes.Equal(hash, repr.Hash), nil
}

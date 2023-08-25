package xutil

import (
	"github.com/42milez/go-oidc-server/app/ent/typedef"
	"github.com/oklog/ulid/v2"
)

func IsEmpty[T string | []byte](v T) bool {
	return len(v) == 0
}

func MakeUserID() typedef.UserID {
	return typedef.UserID(ulid.Make().String())
}

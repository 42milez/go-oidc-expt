package xutil

import (
	"time"

	"github.com/42milez/go-oidc-server/app/idp/ent/typedef"
	"github.com/oklog/ulid/v2"
)

type Clocker interface {
	Now() time.Time
}

type RealClocker struct{}

func (v RealClocker) Now() time.Time {
	return time.Now()
}

func IsEmpty[T string | []byte](v T) bool {
	return len(v) == 0
}

func MakeUserID() typedef.UserID {
	return typedef.UserID(ulid.Make().String())
}

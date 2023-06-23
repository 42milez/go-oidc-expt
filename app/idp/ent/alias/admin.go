package alias

import (
	"github.com/oklog/ulid/v2"
)

type AdminID string

func (v AdminID) IsZero() bool {
	return len(v) == 0
}

func MakeAdminID() AdminID {
	return AdminID(ulid.Make().String())
}

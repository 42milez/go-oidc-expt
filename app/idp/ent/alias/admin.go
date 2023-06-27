package alias

import (
	"github.com/oklog/ulid/v2"
)

type AdminID string

func (v AdminID) IsZero() bool {
	return len(v) == 0
}

func (v AdminID) MarshalBinary() ([]byte, error) {
	return []byte(v), nil
}

func MakeAdminID() AdminID {
	return AdminID(ulid.Make().String())
}

package alias

import (
	"database/sql/driver"

	"github.com/oklog/ulid/v2"
)

type AdminID ulid.ULID

func (v AdminID) IsZero() bool {
	return v.IsZero()
}

func (v AdminID) MarshalBinary() ([]byte, error) {
	return v[:], nil
}

func (v AdminID) Scan(src any) error {
	return v.Scan(src)
}

func (v AdminID) Value() (driver.Value, error) {
	return v.Value()
}

func MakeAdminID() AdminID {
	return AdminID(ulid.Make())
}

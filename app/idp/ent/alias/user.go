package alias

import (
	"database/sql/driver"
	"github.com/oklog/ulid/v2"
)

type UserID ulid.ULID

func (v UserID) Value() (driver.Value, error) {
	return v.Value()
}

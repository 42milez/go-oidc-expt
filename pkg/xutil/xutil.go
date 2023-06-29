package xutil

import (
	"net/http"
	"time"

	"github.com/42milez/go-oidc-server/app/idp/ent/typedef"
	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
)

type Clocker interface {
	Now() time.Time
}

type RealClocker struct{}

func (v RealClocker) Now() time.Time {
	return time.Now()
}

func CloseHTTPConn(resp *http.Response) {
	if err := resp.Body.Close(); err != nil {
		log.Error().Err(err).Msg("failed to close http connection")
	}
}

func IsEmpty[T string | []byte](v T) bool {
	return len(v) == 0
}

func MakeAdminID() typedef.AdminID {
	return typedef.AdminID(ulid.Make().String())
}

package xerr

import (
	"fmt"
)

var (
	FailedToCreateSession    InternalErr = "session id already exists"
	FailedToInitialize       InternalErr = "failed to initialize"
	FailedToReadResponseBody InternalErr = "failed to read response body"
	FailedToValidate         InternalErr = "failed to validate"
	NotFound                 InternalErr = "not found"
	OIDCInvalidResponseType  InternalErr = "invalid response type"
	PasswordNotMatched       InternalErr = "password not matched"
	SessionNotFound          InternalErr = "session not found"
	UnknownSecurityScheme    InternalErr = "unknown security scheme"
	UserNotFound             InternalErr = "user not found"
	AuthCodeNotFound         InternalErr = "auth code not found"
	AuthCodeExpired          InternalErr = "auth code expired"
	AuthCodeUsed             InternalErr = "auth code used"
	MalformedFormParameter   InternalErr = "malformed form parameter"
	RedirectUriNotFound      InternalErr = "redirect uri not found"
	ConsentNotFound          InternalErr = "consent not found"
)

var (
	InvalidRequest              PublicErr = "invalid request"
	InvalidUsernameOrPassword   PublicErr = "invalid username or password"
	OK                          PublicErr = "ok"
	ServiceTemporaryUnavailable PublicErr = "service temporary unavailable"
	UnauthorizedRequest         PublicErr = "unauthorized request"
	UnexpectedErrorOccurred     PublicErr = "unexpected error occurred"
)

type InternalErr string

func (v InternalErr) Error() string {
	return string(v)
}

func (v InternalErr) Wrap(e error) error {
	return fmt.Errorf("%w ( %w )", v, e)
}

type PublicErr string

func (v PublicErr) Error() string {
	return string(v)
}

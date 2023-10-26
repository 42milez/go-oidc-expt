package xerr

import (
	"fmt"
)

var (
	AuthCodeExpired          InternalErr = "auth code expired"
	AuthCodeNotFound         InternalErr = "auth code not found"
	AuthCodeUsed             InternalErr = "auth code used"
	ConsentNotFound          InternalErr = "consent not found"
	ContextValueNotFound     InternalErr = "context value not found"
	FailedToInitialize       InternalErr = "failed to initialize"
	FailedToReadResponseBody InternalErr = "failed to read response body"
	FailedToValidate         InternalErr = "failed to validate"
	FailedToWriteSession     InternalErr = "session id already exists"
	InvalidToken             InternalErr = "invalid refresh token"
	MalformedFormParameter   InternalErr = "malformed form parameter"
	OIDCInvalidResponseType  InternalErr = "invalid response type"
	PasswordNotMatched       InternalErr = "password not matched"
	RedirectUriNotFound      InternalErr = "redirect uri not found"
	RedirectUriNotMatched    InternalErr = "redirect uri not matched"
	SessionNotFound          InternalErr = "session not found"
	UnknownSecurityScheme    InternalErr = "unknown security scheme"
	UserNotFound             InternalErr = "user not found"
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

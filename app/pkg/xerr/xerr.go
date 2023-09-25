package xerr

import (
	"fmt"
)

var (
	FailedToInitialize       InternalErr = "failed to initialize"
	FailedToReadResponseBody InternalErr = "failed to read response body"
	FailedToUnmarshalJSON    InternalErr = "failed to unmarshal json"
	NotFound                 InternalErr = "not found"
	OIDCInvalidResponseType  InternalErr = "invalid response type"
	PasswordNotMatched       InternalErr = "password not matched"
	SessionIDAlreadyExists   InternalErr = "session id already exists"
	SessionNotFound          InternalErr = "session not found"
	UserNotFound             InternalErr = "user not found"
	FailedToValidate         InternalErr = "failed to validate"
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

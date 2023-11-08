package xerr

import (
	"fmt"
)

var (
	AuthCodeExpired                InternalErr = "auth code expired"
	AuthCodeNotFound               InternalErr = "auth code not found"
	AuthCodeUsed                   InternalErr = "auth code used"
	CacheKeyDuplicated             InternalErr = "cache key duplicated"
	CacheKeyNotFound               InternalErr = "cache key not found"
	ClientIdNotMatched             InternalErr = "client id not matched"
	ConsentNotFound                InternalErr = "consent not found"
	FailedToInitialize             InternalErr = "failed to initialize"
	FailedToReadResponseBody       InternalErr = "failed to read response body"
	FailedToSetTimeoutOnCacheKey   InternalErr = "failed to set timeout on cache key"
	FailedToValidate               InternalErr = "failed to validate"
	FailedToWriteCache             InternalErr = "failed to write cache"
	InvalidPath                    InternalErr = "invalid path"
	InvalidResponseType            InternalErr = "invalid response type"
	InvalidToken                   InternalErr = "invalid token"
	MalformedFormParameter         InternalErr = "malformed form parameter"
	PasswordNotMatched             InternalErr = "password not matched"
	RedirectUriNotFound            InternalErr = "redirect uri not found"
	RefreshTokenPermissionNotFound InternalErr = "refresh token permission not found"
	UnknownSecurityScheme          InternalErr = "unknown security scheme"
	UserIdNotFoundInContext        InternalErr = "user id not found in context"
	UserNotFound                   InternalErr = "user not found"
)

var (
	InvalidRequest              PublicErr = "invalid request"
	InvalidUsernameOrPassword   PublicErr = "invalid username or password"
	OK                          PublicErr = "ok"
	ServiceTemporaryUnavailable PublicErr = "service temporary unavailable"
	UnauthorizedRequest         PublicErr = "unauthorized request"
	UnexpectedErrorOccurred     PublicErr = "unexpected error occurred"
)

var (
	TypeAssertionFailed TestErr = "type assertion failed"
)

type InternalErr string

func (ie InternalErr) Error() string {
	return string(ie)
}

func (ie InternalErr) Wrap(e error) error {
	return fmt.Errorf("%w ( %w )", ie, e)
}

type PublicErr string

func (pe PublicErr) Error() string {
	return string(pe)
}

type TestErr string

func (te TestErr) Error() string {
	return string(te)
}

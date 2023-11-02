package xerr

import (
	"fmt"
)

var (
	AuthCodeExpired               InternalErr = "auth code expired"
	AuthCodeNotFound              InternalErr = "auth code not found"
	AuthCodeUsed                  InternalErr = "auth code used"
	ConsentNotFound               InternalErr = "consent not found"
	FailedToInitialize            InternalErr = "failed to initialize"
	FailedToReadResponseBody      InternalErr = "failed to read response body"
	FailedToValidate              InternalErr = "failed to validate"
	FailedToWriteCache            InternalErr = "failed to write cache"
	InvalidResponseType           InternalErr = "invalid response type"
	InvalidToken                  InternalErr = "invalid token"
	MalformedFormParameter        InternalErr = "malformed form parameter"
	PasswordNotMatched            InternalErr = "password not matched"
	RedirectUriNotFound           InternalErr = "redirect uri not found"
	UserIdNotFoundInContext       InternalErr = "user id not found in context"
	UnknownSecurityScheme         InternalErr = "unknown security scheme"
	UserNotFound                  InternalErr = "user not found"
	CacheKeyDuplicated            InternalErr = "cache key duplicated"
	CacheKeyNotFound              InternalErr = "cache key not found"
	FailedToSetTimeoutOnCacheKey  InternalErr = "failed to set timeout on cache key"
	RefreshTokenOwnerIdNotMatched InternalErr = "refresh token owner id not matched"
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

package xerr

import (
	"fmt"
)

var (
	AuthCodeExpired              InternalError = "auth code expired"
	AuthCodeNotFound             InternalError = "auth code not found"
	AuthCodeUsed                 InternalError = "auth code used"
	CacheFieldDuplicated         InternalError = "cache field duplicated"
	CacheKeyNotFound             InternalError = "cache key not found"
	ClientIDNotMatched           InternalError = "client id not matched"
	ConsentNotFound              InternalError = "consent not found"
	CredentialNotFoundInHeader   InternalError = "credential not found in header"
	FailedToInitialize           InternalError = "failed to initialize"
	FailedToReadResponseBody     InternalError = "failed to read response body"
	FailedToReadSession          InternalError = "failed to read session"
	FailedToSetTimeoutOnCacheKey InternalError = "failed to set timeout on cache key"
	FailedToValidate             InternalError = "failed to validate"
	FailedToWriteCache           InternalError = "failed to write cache"
	InvalidPath                  InternalError = "invalid path"
	InvalidResponseType          InternalError = "invalid response type"
	InvalidToken                 InternalError = "invalid token"
	PasswordNotMatched           InternalError = "password not matched"
	RecordNotFound               InternalError = "record not found"
	RedirectUriNotFound          InternalError = "redirect uri not found"
	RefreshTokenNotFound         InternalError = "refresh token not found"
	RefreshTokenNotMatched       InternalError = "refresh token not matched"
	UnauthorizedRequest          InternalError = "unauthorized request"
	UnexpectedErrorOccurred      InternalError = "unexpected error occurred"
	UnknownSecurityScheme        InternalError = "unknown security scheme"
	UserIDNotFoundInContext      InternalError = "user id not found in context"
	UserNotFound                 InternalError = "user not found"
	InvalidRedirectURI           InternalError = "invalid redirect uri"
	UnsupportedClaim             InternalError = "unsupported claim"
)

var (
	InvalidRequest2             PublicError = "invalid request"
	InvalidUsernameOrPassword   PublicError = "invalid username or password"
	OK                          PublicError = "ok"
	ServiceTemporaryUnavailable PublicError = "service temporary unavailable"
	UnexpectedErrorOccurred2    PublicError = "unexpected error occurred"
)

var (
	AccessDenied   OIDCError = "access_denied"
	InvalidClient  OIDCError = "invalid_client"
	InvalidGrant   OIDCError = "invalid_grant"
	InvalidRequest OIDCError = "invalid_request"
	ServerError    OIDCError = "server_error"
)

var (
	TypeAssertionFailed TestError = "type assertion failed"
)

type InternalError string

func (ie InternalError) Error() string {
	return string(ie)
}

func (ie InternalError) Wrap(e error) error {
	return fmt.Errorf("%w ( %w )", ie, e)
}

type PublicError string

func (pe PublicError) Error() string {
	return string(pe)
}

type OIDCError string

func (oe OIDCError) Error() string {
	return string(oe)
}

type TestError string

func (te TestError) Error() string {
	return string(te)
}

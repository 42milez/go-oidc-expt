package xerr

import (
	"fmt"
)

var (
	AuthCodeExpired                InternalError = "auth code expired"
	AuthCodeNotFound               InternalError = "auth code not found"
	AuthCodeUsed                   InternalError = "auth code used"
	CacheKeyDuplicated             InternalError = "cache key duplicated"
	CacheKeyNotFound               InternalError = "cache key not found"
	ClientIdNotMatched             InternalError = "client id not matched"
	ConsentNotFound                InternalError = "consent not found"
	FailedToInitialize             InternalError = "failed to initialize"
	FailedToReadResponseBody       InternalError = "failed to read response body"
	FailedToSetTimeoutOnCacheKey   InternalError = "failed to set timeout on cache key"
	FailedToValidate               InternalError = "failed to validate"
	FailedToWriteCache             InternalError = "failed to write cache"
	InvalidPath                    InternalError = "invalid path"
	InvalidResponseType            InternalError = "invalid response type"
	InvalidToken                   InternalError = "invalid token"
	MalformedFormParameter         InternalError = "malformed form parameter"
	PasswordNotMatched             InternalError = "password not matched"
	RedirectUriNotFound            InternalError = "redirect uri not found"
	RefreshTokenPermissionNotFound InternalError = "refresh token permission not found"
	UnknownSecurityScheme          InternalError = "unknown security scheme"
	UserIdNotFoundInContext        InternalError = "user id not found in context"
	UserNotFound                   InternalError = "user not found"
	CredentialNotFoundInHeader     InternalError = "credential not found in header"
	UnauthorizedRequest            InternalError = "unauthorized request"
	ParameterNotFound              InternalError = "parameter not found"
	RedirectURINotMatched          InternalError = "redirect uri not matched"
	RecordNotFound                 InternalError = "record not found"
	FailedToRevokeAuthCode         InternalError = "failed to revoke authorization code"
)

var (
	InvalidRequest2             PublicError = "invalid request"
	InvalidUsernameOrPassword   PublicError = "invalid username or password"
	OK                          PublicError = "ok"
	ServiceTemporaryUnavailable PublicError = "service temporary unavailable"
	UnexpectedErrorOccurred     PublicError = "unexpected error occurred"
)

var (
	AccessDenied            OIDCError = "access_denied"
	InvalidClient           OIDCError = "invalid_client"
	InvalidGrant            OIDCError = "invalid_grant"
	InvalidRequest          OIDCError = "invalid_request"
	InvalidScope            OIDCError = "invalid_scope"
	ServerError             OIDCError = "server_error"
	TemporaryUnavailable    OIDCError = "temporarily_unavailable"
	UnauthorizedClient      OIDCError = "unauthorized_client"
	UnsupportedGrantType    OIDCError = "unsupported_grant_type"
	UnsupportedResponseType OIDCError = "unsupported_response_type"
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

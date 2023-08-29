package xerr

import (
	"errors"
	"fmt"
)

type OpenIDErr error

var (
	InvalidResponseType = errors.New("invalid response type")
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

// --------------------------------------------------

var (
	AuthenticationFailed        PublicErr = "authentication failed"
	InvalidParameter            PublicErr = "invalid parameter"
	InvalidRequest              PublicErr = "invalid request"
	InvalidUsernameOrPassword   PublicErr = "invalid username or password"
	OK                          PublicErr = "ok"
	ServiceCurrentlyUnavailable PublicErr = "service currently unavailable"
	UnauthorizedUser            PublicErr = "unauthorized user"
	UnexpectedErrorOccurred     PublicErr = "unexpected error occurred"
)

var (
	FailToEstablishConnection    InternalErr = "failed to establish connection"
	FailedToBuildToken           InternalErr = "failed to build token"
	FailedToCloseConnection      InternalErr = "failed to close connection"
	FailedToCloseResponseBody    InternalErr = "failed to close response body"
	FailedToCreateUser           InternalErr = "failed to create user"
	FailedToDecodeInToBytes      InternalErr = "failed to decode string"
	FailedToDecodeInToStruct     InternalErr = "failed to decode into struct"
	FailedToDeleteItem           InternalErr = "failed to delete item"
	FailedToEncodeInToBytes      InternalErr = "failed to encode struct"
	FailedToExtractToken         InternalErr = "failed to extract token"
	FailedToGenerateRandomBytes  InternalErr = "failed to generate random bytes"
	FailedToGenerateUniqueID     InternalErr = "failed to generate unique id"
	FailedToHashPassword         InternalErr = "failed to hash password"
	FailedToInitialize           InternalErr = "failed to initialize"
	FailedToLoadItem             InternalErr = "failed to load item"
	FailedToParseEnvVal          InternalErr = "failed to parse environment variable"
	FailedToParsePrivateKey      InternalErr = "failed to parse private key"
	FailedToParsePublicKey       InternalErr = "failed to parse public key"
	FailedToParseRequest         InternalErr = "failed to parse request"
	FailedToPingCache            InternalErr = "failed to ping cache"
	FailedToPingDatabase         InternalErr = "failed to ping database"
	FailedToReachHost            InternalErr = "failed to reach host"
	FailedToReadContextValue     InternalErr = "failed to read context value"
	FailedToReadFile             InternalErr = "failed to read file"
	FailedToReadResponseBody     InternalErr = "failed to read response body"
	FailedToReadResponseLocation InternalErr = "failed to read response location"
	FailedToSetInToCache         InternalErr = "failed to set into cache"
	FailedToSignToken            InternalErr = "failed to sign token"
	FailedToUnmarshalJSON        InternalErr = "failed to unmarshal json"
	InvalidToken                 InternalErr = "invalid token"
	PasswordNotMatched           InternalErr = "password not matched"
	ResponseBodyNotMatched       InternalErr = "response body not matched"
	SessionIDAlreadyExists       InternalErr = "session id already exists"
	UserNotFound                 InternalErr = "user not found"
)

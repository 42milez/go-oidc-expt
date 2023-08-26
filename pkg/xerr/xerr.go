package xerr

import (
	"errors"
	"fmt"
)

const (
	FailedToCloseConnection      GeneralErr = "failed to close connection"
	FailedToCloseResponseBody    GeneralErr = "failed to close response body"
	FailedToInitialize           GeneralErr = "failed to initialize"
	FailedToPingCache            GeneralErr = "failed to ping cache"
	FailedToPingDatabase         GeneralErr = "failed to ping database"
	FailedToReadContextValue     GeneralErr = "failed to read context value"
	FailedToReadFile             GeneralErr = "failed to read file"
	FailedToReadResponseBody     GeneralErr = "failed to read response body"
	FailedToReadResponseLocation GeneralErr = "failed to read response location"
	FailedToUnmarshalJSON        GeneralErr = "failed to unmarshal json"
)

type GeneralErr string

func (v GeneralErr) Error() string {
	return string(v)
}

type HTTPErr string

func (v HTTPErr) Error() string {
	return string(v)
}

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
	FailToEstablishConnection   InternalErr = "failed to establish connection"
	FailedToBuildToken          InternalErr = "failed to build token"
	FailedToCreateUser          InternalErr = "failed to create user"
	FailedToDecodeInToBytes     InternalErr = "failed to decode string"
	FailedToDecodeInToStruct    InternalErr = "failed to decode into struct"
	FailedToDeleteItem          InternalErr = "failed to delete item"
	FailedToEncodeInToBytes     InternalErr = "failed to encode struct"
	FailedToExtractToken        InternalErr = "failed to extract token"
	FailedToGenerateRandomBytes InternalErr = "failed to generate random bytes"
	FailedToHashPassword        InternalErr = "failed to hash password"
	FailedToLoadItem            InternalErr = "failed to load item"
	FailedToParseEnvVal         InternalErr = "failed to parse environment variable"
	FailedToParsePrivateKey     InternalErr = "failed to parse private key"
	FailedToParsePublicKey      InternalErr = "failed to parse public key"
	FailedToParseRequest        InternalErr = "failed to parse request"
	FailedToReachHost           InternalErr = "failed to reach host"
	FailedToSaveItem            InternalErr = "failed to save item"
	FailedToSignToken           InternalErr = "failed to sign token"
	InvalidToken                InternalErr = "invalid token"
	PasswordNotMatched          InternalErr = "password not matched"
)

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

const (
	AuthenticationFailed        HTTPErr = "authentication failed"
	InvalidParameter            HTTPErr = "invalid parameter"
	InvalidRequest              HTTPErr = "invalid request"
	ServiceCurrentlyUnavailable HTTPErr = "service currently unavailable"
	UnauthorizedUser            HTTPErr = "unauthorized user"
	UnexpectedErrorOccurred     HTTPErr = "unexpected error occurred"
)

type HTTPErr string

func (v HTTPErr) Error() string {
	return string(v)
}

type OpenIDErr error

var (
	InvalidResponseType = errors.New("invalid response type")
)

type Err string

func (v Err) Error() string {
	return string(v)
}

func (v Err) Wrap(e error) error {
	return fmt.Errorf("%w ( %w )", v, e)
}

// --------------------------------------------------

var (
	FailToEstablishConnection   Err = "failed to establish connection"
	FailedToBuildToken          Err = "failed to build token"
	FailedToCreateUser          Err = "failed to create user"
	FailedToDecodeInToBytes     Err = "failed to decode string"
	FailedToDecodeInToStruct    Err = "failed to decode into struct"
	FailedToDeleteItem          Err = "failed to delete item"
	FailedToEncodeInToBytes     Err = "failed to encode struct"
	FailedToExtractToken        Err = "failed to extract token"
	FailedToGenerateRandomBytes Err = "failed to generate random bytes"
	FailedToHashPassword        Err = "failed to hash password"
	FailedToLoadItem            Err = "failed to load item"
	FailedToParseEnvVal         Err = "failed to parse environment variable"
	FailedToParsePrivateKey     Err = "failed to parse private key"
	FailedToParsePublicKey      Err = "failed to parse public key"
	FailedToParseRequest        Err = "failed to parse request"
	FailedToReachHost           Err = "failed to reach host"
	FailedToSaveItem            Err = "failed to save item"
	FailedToSignToken           Err = "failed to sign token"
	InvalidToken                Err = "invalid token"
	OK                          Err = "ok"
	PasswordNotMatched          Err = "password not matched"
)

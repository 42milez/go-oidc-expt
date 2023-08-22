package xerr

import (
	"errors"
	"fmt"
)

const (
	FailToEstablishConnection    GeneralErr = "failed to establish connection"
	FailedToCloseConnection      GeneralErr = "failed to close connection"
	FailedToCloseResponseBody    GeneralErr = "failed to close response body"
	FailedToDecodeInToBytes      GeneralErr = "failed to decode string"
	FailedToEncodeInToBytes      GeneralErr = "failed to encode struct"
	FailedToGenerateRandomBytes  GeneralErr = "failed to generate random bytes"
	FailedToInitialize           GeneralErr = "failed to initialize"
	FailedToPingCache            GeneralErr = "failed to ping cache"
	FailedToPingDatabase         GeneralErr = "failed to ping database"
	FailedToReachHost            GeneralErr = "failed to reach host"
	FailedToReadContextValue     GeneralErr = "failed to read context value"
	FailedToReadFile             GeneralErr = "failed to read file"
	FailedToReadResponseBody     GeneralErr = "failed to read response body"
	FailedToReadResponseLocation GeneralErr = "failed to read response location"
	FailedToUnmarshalJSON        GeneralErr = "failed to unmarshal json"
	FailedToDecodeInToStruct     GeneralErr = "failed to decode into struct"
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

func Wrap(e1, e2 error) error {
	return fmt.Errorf("%w (%w)", e1, e2)
}

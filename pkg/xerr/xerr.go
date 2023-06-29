package xerr

import "fmt"

const (
	FailedToCloseConnection   GeneralErr = "failed to close connection"
	FailedToCloseResponseBody GeneralErr = "failed to close response body"
	FailToEstablishConnection GeneralErr = "failed to establish connection"
	FailedToInitialize        GeneralErr = "failed to initialize"
	FailedToPingCache         GeneralErr = "failed to ping cache"
	FailedToPingDatabase      GeneralErr = "failed to ping database"
	FailedToReachHost         GeneralErr = "failed to reach host"
	FailedToReadContextValue  GeneralErr = "failed to read context value"
	FailedToReadFile          GeneralErr = "failed to read file"
	FailedToResponseBody      GeneralErr = "failed to response body"
	FailedToUnmarshalJSON     GeneralErr = "failed to unmarshal json"
	UnexpectedValue           GeneralErr = "unexpected value"
)

type GeneralErr string

func (v GeneralErr) Error() string {
	return string(v)
}

const (
	AuthenticationFailed        HTTPErr = "authentication failed"
	ServiceCurrentlyUnavailable HTTPErr = "service currently unavailable"
	UnexpectedErrorOccurred     HTTPErr = "unexpected error occurred"
)

type HTTPErr string

func (v HTTPErr) Error() string {
	return string(v)
}

func WrapErr(e1, e2 error) error {
	return fmt.Errorf("%w:%w", e1, e2)
}

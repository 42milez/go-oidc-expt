package xerr

const (
	FailedToCloseConnection   GeneralErr = "failed to close connection"
	FailedToCloseResponseBody GeneralErr = "failed to close response body"
	FailedToInitialize        GeneralErr = "failed to initialize"
	FailedToReachHost         GeneralErr = "failed to reach host"
	FailedToReadContextValue  GeneralErr = "failed to read context value"
	FailedToReadFile          GeneralErr = "failed to read file"
	FailedToResponseBody      GeneralErr = "failed to response body"
	FailedToUnmarshalJSON     GeneralErr = "failed to unmarshal json"
)

type GeneralErr string

func (v GeneralErr) Error() string {
	return string(v)
}

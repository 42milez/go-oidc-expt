package xerr

const (
	AuthenticationFailed        HTTPErr = "authentication failed"
	ServiceCurrentlyUnavailable HTTPErr = "service currently unavailable"
	UnexpectedErrorOccurred     HTTPErr = "unexpected error occurred"
)

type HTTPErr string

func (v HTTPErr) Error() string {
	return string(v)
}

package xerr

const (
	OK HTTPErr = "ok"
	ErrFailedToAuthenticate HTTPErr = "failed to authenticate"
	ErrInternalServerError HTTPErr = "internal server error"
)

type HTTPErr string

func (v HTTPErr) Error() string {
	return string(v)
}

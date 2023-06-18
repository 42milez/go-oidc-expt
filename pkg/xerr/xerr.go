package xerr

const (
	FailedToCloseConnection GeneralErr = "failed to close connection"
	FailedToInitialize      GeneralErr = "failed to initialize"
	FailedToReachHost       GeneralErr = "failed to reach host"
)

type GeneralErr string

func (v GeneralErr) Error() string {
	return string(v)
}

package testutil

const DummyError DummyErr = "DUMMY ERROR"

type DummyErr string

func (v DummyErr) Error() string {
	return string(v)
}

package util

import "time"

type Clocker interface {
	Now() time.Time
}

type RealClocker struct{}

func (v RealClocker) Now() time.Time {
	return time.Now()
}

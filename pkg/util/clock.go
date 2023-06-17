package util

import "time"

type Clocker interface {
	Now() time.Time
}

type RealClocker struct{}

func (v RealClocker) Now() time.Time {
	return time.Now()
}

type FixedClocker struct{}

func (v FixedClocker) Now() time.Time {
	return time.Date(2000, 12, 31, 23, 59, 59, 0, time.UTC)
}

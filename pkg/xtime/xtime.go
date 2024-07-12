package xtime

import "time"

type RealClocker struct{}

func (v RealClocker) Now() time.Time {
	return time.Now().UTC()
}

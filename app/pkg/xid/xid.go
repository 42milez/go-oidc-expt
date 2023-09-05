package xid

import (
	"github.com/sony/sonyflake"
)

var UID *UniqueID

// TODO: Separate the ID generator as an independent container.

func Init() error {
	sf, err := sonyflake.New(sonyflake.Settings{})
	if err != nil {
		return err
	}
	UID = &UniqueID{
		generator: sf,
	}
	return nil
}

type UniqueID struct {
	generator *sonyflake.Sonyflake
}

func (p *UniqueID) NextID() (uint64, error) {
	return p.generator.NextID()
}

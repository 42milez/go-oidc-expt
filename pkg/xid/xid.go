package xid

import (
	"github.com/sony/sonyflake"
)

var uid *UniqueID

func GetUniqueIDGenerator() (*UniqueID, error) {
	if uid != nil {
		return uid, nil
	}

	sf, err := sonyflake.New(sonyflake.Settings{})

	if err != nil {
		return nil, err
	}

	uid = &UniqueID{
		generator: sf,
	}

	return uid, nil
}

type UniqueID struct {
	generator *sonyflake.Sonyflake
}

func (p *UniqueID) NextID() (uint64, error) {
	return p.generator.NextID()
}

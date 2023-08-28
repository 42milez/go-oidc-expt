package xid

import "github.com/sony/sonyflake"

type UniqueID struct {
	generator *sonyflake.Sonyflake
}

func (p *UniqueID) NextID() (uint64, error) {
	return p.NextID()
}

func NewUniqueID() (*UniqueID, error) {
	sf, err := sonyflake.New(sonyflake.Settings{})
	return &UniqueID{
		generator: sf,
	}, err
}

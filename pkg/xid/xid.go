package xid

import (
	"log"

	"github.com/sony/sonyflake"
)

var UID *UniqueID

// TODO: Separate the ID generator as an independent container.

type UniqueID struct {
	generator *sonyflake.Sonyflake
}

func (p *UniqueID) NextID() (uint64, error) {
	return p.generator.NextID()
}

func init() {
	sf, err := sonyflake.New(sonyflake.Settings{})
	if err != nil {
		log.Fatal(err)
	}
	UID = &UniqueID{
		generator: sf,
	}
}

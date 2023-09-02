package repository

type IDGenerator interface {
	NextID() (uint64, error)
}

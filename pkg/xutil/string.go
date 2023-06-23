package xutil

func IsEmpty[T string | []byte](v T) bool {
	return len(v) == 0
}

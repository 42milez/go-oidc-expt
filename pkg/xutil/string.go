package xutil

func ComparePassword(pw1, pw2 string) error {
	return nil
}

func IsEmpty[T string | []byte](v T) bool {
	return len(v) == 0
}

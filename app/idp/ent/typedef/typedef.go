package typedef

type UserID string

func (v UserID) IsZero() bool {
	return len(v) == 0
}

func (v UserID) MarshalBinary() ([]byte, error) {
	return []byte(v), nil
}

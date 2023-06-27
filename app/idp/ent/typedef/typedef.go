package typedef

type PasswordHash string

type AdminID string

func (v AdminID) IsZero() bool {
	return len(v) == 0
}

func (v AdminID) MarshalBinary() ([]byte, error) {
	return []byte(v), nil
}

type UserID string

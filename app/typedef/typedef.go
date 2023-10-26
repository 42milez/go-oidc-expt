package typedef

import "strconv"

type AuthCodeID uint64
type ConsentID uint64
type RedirectUriID uint64
type RelyingPartyID uint64
type SessionID uint64
type TokenType string
type UserID uint64

type SessionIdKey struct{}
type UserIdKey struct{}

func (sid SessionID) String() string {
	return strconv.FormatUint(uint64(sid), 10)
}

func (uid UserID) MarshalBinary() ([]byte, error) {
	return []byte(strconv.FormatUint(uint64(uid), 10)), nil
}

func (uid UserID) String() string {
	return strconv.FormatUint(uint64(uid), 10)
}

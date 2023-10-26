package typedef

import "strconv"

type AuthCodeID uint64
type ConsentID uint64
type RelyingPartyID uint64
type RedirectUriID uint64
type TokenType string

type SessionIdKey struct{}
type UserIdKey struct{}

type SessionID uint64

func (sid SessionID) String() string {
	return strconv.FormatUint(uint64(sid), 10)
}

type UserID uint64

func (uid UserID) MarshalBinary() ([]byte, error) {
	return []byte(strconv.FormatUint(uint64(uid), 10)), nil
}

func (uid UserID) String() string {
	return strconv.FormatUint(uint64(uid), 10)
}

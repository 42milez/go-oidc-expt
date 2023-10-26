package typedef

import "strconv"

type AuthCodeID uint64
type ConsentID uint64
type RelyingPartyID uint64
type RedirectUriID uint64
type SessionID uint64
type UserID uint64
type TokenType string

type SessionIdKey struct{}
type UserIdKey struct{}

func (uid UserID) MarshalBinary() ([]byte, error) {
	return []byte(strconv.FormatUint(uint64(uid), 10)), nil
}

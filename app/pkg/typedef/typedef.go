package typedef

import "strconv"

//  Entity
// --------------------------------------------------

type AuthCodeID uint64
type ConsentID uint64
type RedirectURIID uint64
type RelyingPartyID uint64
type SessionID uint64
type UserID uint64

//  Context
// --------------------------------------------------

type RequestParamKey struct{}

type SessionIdKey struct{}

func (sid SessionID) String() string {
	return strconv.FormatUint(uint64(sid), 10)
}

type UserIdKey struct{}

func (uid UserID) MarshalBinary() ([]byte, error) {
	return []byte(strconv.FormatUint(uint64(uid), 10)), nil
}

func (uid UserID) String() string {
	return strconv.FormatUint(uint64(uid), 10)
}

//  Cache
// --------------------------------------------------

type OpenIdParam struct {
	RedirectURI string
	UserId      UserID
}

type RefreshTokenPermission struct {
	ClientId string
	UserId   UserID
}

//  API
// --------------------------------------------------

type TokenType string
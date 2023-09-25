package service

import (
	"github.com/42milez/go-oidc-server/app/pkg/xtime"
	"github.com/gorilla/securecookie"
)

//  AUTHENTICATION
// --------------------------------------------------

func NewAuthenticate(repo UserConsentReader, tokenGen TokenGenerator) *Authenticate {
	return &Authenticate{
		repo:     repo,
		tokenGen: tokenGen,
	}
}

//  OIDC: Authorization
// --------------------------------------------------

func NewAuthorize(repo Authorizer) *Authorize {
	return &Authorize{
		repo: repo,
	}
}

func NewConsent(repo ConsentCreator) *AcceptConsent {
	return &AcceptConsent{
		repo: repo,
	}
}

//  COOKIE
// --------------------------------------------------

func NewCookie(hashKey, blockKey []byte, clock xtime.Clocker) *Cookie {
	return &Cookie{
		clock: clock,
		sc:    securecookie.New(hashKey, blockKey),
	}
}

//  HEALTH CHECK
// --------------------------------------------------

func NewCheckHealth(repo HealthChecker) *CheckHealth {
	return &CheckHealth{
		repo: repo,
	}
}

//  SESSION
// --------------------------------------------------

func NewCreateSession(repo SessionCreator) *CreateSession {
	return &CreateSession{
		repo: repo,
	}
}

func NewRestoreSession(repo SessionReader) *RestoreSession {
	return &RestoreSession{
		repo: repo,
	}
}

func NewUpdateSession(repo SessionUpdater) *UpdateSession {
	return &UpdateSession{
		repo: repo,
	}
}

//  USER
// --------------------------------------------------

func NewCreateUser(repo UserCreator) *CreateUser {
	return &CreateUser{
		repo: repo,
	}
}

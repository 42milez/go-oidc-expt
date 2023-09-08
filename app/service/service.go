package service

import (
	"github.com/42milez/go-oidc-server/app/datastore"
	"github.com/42milez/go-oidc-server/app/repository"
)

// --------------------------------------------------
//  AUTHENTICATION
// --------------------------------------------------

func NewAuthenticate(repo UserSelector, tokenGen TokenGenerator) *Authenticate {
	return &Authenticate{
		repo:     repo,
		tokenGen: tokenGen,
	}
}

// --------------------------------------------------
//  HEALTH CHECK
// --------------------------------------------------

func NewCheckHealth(repo HealthChecker) *CheckHealth {
	return &CheckHealth{
		repo: repo,
	}
}

// --------------------------------------------------
//  SESSION
// --------------------------------------------------

func NewCreateSession(cache *datastore.Cache) *CreateSession {
	return &CreateSession{
		repo: &repository.CreateSession{
			Cache: cache,
		},
	}
}

func NewRestoreSession(cache *datastore.Cache) *RestoreSession {
	return &RestoreSession{
		repo: &repository.ReadSession{
			Cache: cache,
		},
	}
}

// --------------------------------------------------
//  USER
// --------------------------------------------------

func NewCreateUser(repo UserCreator) *CreateUser {
	return &CreateUser{
		repo: repo,
	}
}

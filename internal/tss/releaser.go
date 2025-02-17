package tss

import "github.com/nuvosphere/nudex-voter/internal/types/party"

// Releaser is the interface that wraps the basic Release method.
type Releaser interface {
	// Release releases associated resources. Release should always success
	// and can be called multiple times without causing error.
	Release()
}

type SessionReleaser interface {
	SessionRelease(session party.SessionID)
}

type TODO struct{}

func (t TODO) SessionRelease(session party.SessionID) {}

func (t TODO) Release() {}

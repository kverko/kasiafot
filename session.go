package main

import (
	"fmt"
	"sync"
)

type Manager struct {
	cookieName string
	lock       sync.Mutex
	provider   Provider
	lifetime   int64
}

func NewManager(provideName, cookieName string, lifetime int64) (*Manager, error) {
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("sessions: unknown provide %q", provideName)
	}
	return &Manager{provider: provider, cookieName: cookieName, lifetime: lifetime}
}

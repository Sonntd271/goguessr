package types

import "sync"

type SessionStorage struct {
	sync.RWMutex
	Sessions map[string]bool
}

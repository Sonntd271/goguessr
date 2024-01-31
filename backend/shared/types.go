package shared

import "sync"

type SessionStorage struct {
	sync.RWMutex
	Sessions map[string]bool
}

package api

import (
	"backend/shared"
)

func ValidateToken(token string, ss *shared.SessionStorage) bool {
	// ValidateToken accepts a string and records it as an entry in an in-memory storage, i.e. map
	// Uses Mutex ReadLock for its functionality
	ss.RLock()
	defer ss.RUnlock()

	valid, ok := ss.Sessions[token]
	if !ok {
		return ok
	}
	return valid
}

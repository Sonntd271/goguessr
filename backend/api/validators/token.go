package api

import (
	"backend/shared"
)

func ValidateToken(token string, ss *shared.SessionStorage) bool {
	ss.RLock()
	defer ss.RUnlock()

	valid, ok := ss.Sessions[token]
	if !ok {
		return ok
	}
	return valid
}

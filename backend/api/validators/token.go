package api

import (
	"backend/types"
)

func ValidateToken(token string, ss *types.SessionStorage) bool {
	ss.RLock()
	defer ss.RUnlock()

	valid, ok := ss.Sessions[token]
	if !ok {
		return ok
	}
	return valid
}

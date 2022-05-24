package screenplay

import "backend/user"

var pm *PermissionsManager

func init() {
	pm = new(PermissionsManager)
}

type PermissionsManager struct {
}

func (p *PermissionsManager) UpdateScreenplay(user *user.User, unupdatedScreenplay *Screenplay) (bool, error) {
	if !user.Boss {
		return false, m.err.NoAuthority
	}
	if unupdatedScreenplay.Store != user.Store {
		return false, nil
	}
	return true, nil
}

func GetPermissionsManager() *PermissionsManager {
	return pm
}

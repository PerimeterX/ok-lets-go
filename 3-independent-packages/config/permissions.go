package config

import (
	"time"
)

// Here we declare an interface representing our local
// needs from the configuration db, namely,
// the `AllPermissions` method.
type PermissionDB interface {
	AllPermissions() map[string][]string // maps from role to its permissions
}

// Then we export a service than will provide the
// permissions from memory, to use it, another package
// will have to declare a local interface.
type PermissionStorage struct {
	permissions map[string][]string
}

func NewPermissionStorage(db PermissionDB) *PermissionStorage {
	s := &PermissionStorage{}
	s.permissions = db.AllPermissions()
	go func() {
		for {
			time.Sleep(time.Hour)
			s.permissions = db.AllPermissions()
		}
	}()
	return s
}

func (s *PermissionStorage) RolePermissions(role string) []string {
	return s.permissions[role]
}

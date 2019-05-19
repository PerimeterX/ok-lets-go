package handler

import (
	"fmt"
	"net/http"
	"strings"
)

// Declaring our local needs from the user db instance,
type UserDB interface {
	UserRoleByID(id string) string
}

// ... and our local needs from the in memory permission storage.
type PermissionStorage interface {
	RolePermissions(role string) []string
}

// Lastly, our handler cannot be purely functional,
// since it requires references to non singleton instances.
type UserPermissionsByID struct {
	UserDB             UserDB
	PermissionsStorage PermissionStorage
}

func (u *UserPermissionsByID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query()["id"][0]
	role := u.UserDB.UserRoleByID(id)
	permissions := u.PermissionsStorage.RolePermissions(role)
	fmt.Fprint(w, strings.Join(permissions, ", "))
}

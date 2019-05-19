package main

import (
	"fmt"
	"net/http"
	"strings"
)

func UserPermissionsByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query()["id"][0]
	role := userDBInstance.userRoleByID(id)
	permissions := rolePermissions[role]
	fmt.Fprint(w, strings.Join(permissions, ", "))
}

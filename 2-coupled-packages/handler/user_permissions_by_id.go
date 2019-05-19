package handler

import (
	"fmt"
	"github.com/perimeterx/ok-lets-go/2-coupled-packages/definition"
	"net/http"
	"strings"
)

func UserPermissionsByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query()["id"][0]
	role := definition.UserDBInstance.UserRoleByID(id)
	permissions := definition.RolePermissions[role]
	fmt.Fprint(w, strings.Join(permissions, ", "))
}

package main

// Note how the main package is the only one importing
// other local packages.
import (
	"github.com/perimeterx/ok-lets-go/3-independent-packages/config"
	"github.com/perimeterx/ok-lets-go/3-independent-packages/database"
	"github.com/perimeterx/ok-lets-go/3-independent-packages/handler"
	"net/http"
)

func main() {
	userDB := &database.SomeUserDB{}
	configDB := &database.SomeConfigDB{}
	permissionStorage := config.NewPermissionStorage(configDB)
	h := &handler.UserPermissionsByID{UserDB: userDB, PermissionsStorage: permissionStorage}
	http.Handle("/", h)
	http.ListenAndServe(":8080", nil)
}

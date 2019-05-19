package main

// Note how the main package is the only one importing
// packages other than the definition package.
import (
	"github.com/perimeterx/ok-lets-go/2-coupled-packages/config"
	"github.com/perimeterx/ok-lets-go/2-coupled-packages/database"
	"github.com/perimeterx/ok-lets-go/2-coupled-packages/definition"
	"github.com/perimeterx/ok-lets-go/2-coupled-packages/handler"
	"net/http"
)

func main() {
	// This approach also uses singleton instances, and
	// again it's the initiator's responsibility to make
	// sure they're initialized.
	definition.UserDBInstance = &database.SomeUserDB{}
	definition.ConfigDBInstance = &database.SomeConfigDB{}
	config.InitPermissions()
	http.HandleFunc("/", handler.UserPermissionsByID)
	http.ListenAndServe(":8080", nil)
}

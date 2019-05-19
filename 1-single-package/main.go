package main

import (
	"net/http"
	"time"
)

// As noted above, since we plan to only have one instance
// for those 3 services, we'll declare a singleton instance,
// and make sure we only use them to access those services.
var (
	userDBInstance   userDB
	configDBInstance configDB
	rolePermissions  map[string][]string
)

func main() {
	// Our singleton instances will later be assumed
	// initialized, it is the initiator's responsibility
	// to initialize them.
	// The main function will do it with concrete
	// implementation, and test cases, if we plan to
	// have those, may use mock implementations instead.
	userDBInstance = &someUserDB{}
	configDBInstance = &someConfigDB{}
	initPermissions()
	http.HandleFunc("/", UserPermissionsByID)
	http.ListenAndServe(":8080", nil)
}

// This will keep our permissions up to date in memory.
func initPermissions() {
	rolePermissions = configDBInstance.allPermissions()
	go func() {
		for {
			time.Sleep(time.Hour)
			rolePermissions = configDBInstance.allPermissions()
		}
	}()
}

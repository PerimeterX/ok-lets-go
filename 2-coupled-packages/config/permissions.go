package config

import (
	"github.com/perimeterx/ok-lets-go/2-coupled-packages/definition"
	"time"
)

// Since the definition package must not contain any logic,
// managing configuration is implemented in a config package.
func InitPermissions() {
	definition.RolePermissions = definition.ConfigDBInstance.AllPermissions()
	go func() {
		for {
			time.Sleep(time.Hour)
			definition.RolePermissions = definition.ConfigDBInstance.AllPermissions()
		}
	}()
}

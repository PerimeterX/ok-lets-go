package definition

// Note that in this approach both the singleton instance
// and its interface type are declared in the definition
// package. Make sure this package does not contain any
// logic, otherwise it might need to import other packages
// and its neutral nature is compromised.
var (
	UserDBInstance   UserDB
	ConfigDBInstance ConfigDB
)

type UserDB interface {
	UserRoleByID(id string) string
}

type ConfigDB interface {
	AllPermissions() map[string][]string // maps from role to its permissions
}

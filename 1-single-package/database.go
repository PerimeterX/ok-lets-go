package main

// We use interfaces as the types of our database instances
// to make it possible to write tests and use mock implementations.
type userDB interface {
	userRoleByID(id string) string
}

// Note the naming `someConfigDB`. In actual cases we use
// some DB implementation and name our structs accordingly.
// For example, if we use MongoDB, we name our concrete
// struct `mongoConfigDB`. If used in test cases,
// a `mockConfigDB` can be declared, too.
type someUserDB struct{}

func (db *someUserDB) userRoleByID(id string) string {
	// Omitting the implementation details for clarity...
}

type configDB interface {
	allPermissions() map[string][]string // maps from role to its permissions
}

type someConfigDB struct{}

func (db *someConfigDB) allPermissions() map[string][]string {
	// implementation
}

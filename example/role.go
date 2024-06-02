package example

import "golang.org/x/exp/maps"

const (
	RoleOwner  string = "OWNER"
	RoleMember string = "MEMBER"
)

type Role string

func (Role) Map() map[string]int {
	return map[string]int{
		RoleOwner:  10,
		RoleMember: 20,
	}
}

func (r Role) Values() (kinds []string) {
	return maps.Keys(r.Map())
}

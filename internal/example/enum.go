package example

const (
	RoleOwner  Role = "OWNER"
	RoleMember Role = "MEMBER"
)

type Role string

func (r Role) Values() (kinds []string) {
	return []string{
		string(RoleOwner),
		string(RoleMember),
	}
}

// type Level int

// const (
// 	LevelUnknown Level = (iota * 10)
// 	LevelLow
// 	LevelMid
// 	LevelHigh
// )

package auth

type User struct {
	UserName string
	FullName string
	Role     Roles
}

type Roles string

const (
	ADMIN  Roles = "admin"
	EDITOR Roles = "editor"
	BASIC  Roles = "basic"
	NONE   Roles = "none"
)

func ParseRole(role string) Roles {
	switch role {
	case "admin":
		return ADMIN
	case "editor":
		return EDITOR
	case "basic":
		return BASIC
	default:
		return NONE
	}
}

package spec

type Role string

type Enum map[string]Role

func (r Role) String() string {
	return string(r)
}

const (
	admin = "ADMIN"
	user  = "USER"
)

var Roles = Enum{
	admin: admin,
	user:  user,
}

func (e Enum) Admin() Role {
	return e[admin]
}

func (e Enum) User() Role {
	return e[user]
}

func (e Enum) Of(code string) (Role, error) {
	role, ok := e[code]
	if !ok {
		return "", ErrUnsupportedRole(code)
	}

	return role, nil
}

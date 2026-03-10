package app

type User struct {
	Name string
	Role string
}

func NewUser(name string, role string) User {
	return User{
		Name: name,
		Role: role,
	}
}

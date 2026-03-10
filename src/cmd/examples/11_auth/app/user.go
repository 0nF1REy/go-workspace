package app

type User struct {
	Username string
	Password string
}

func NewUser(username, password string) User {
	return User{
		Username: username,
		Password: password,
	}
}

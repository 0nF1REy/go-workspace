package app

type Session struct {
	Username string
	Active   bool
}

func NewSession(username string) Session {
	return Session{
		Username: username,
		Active:   true,
	}
}

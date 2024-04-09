package model

type User struct {
	username string
	password string
}

func NewUser(username string, password string) *User {
	return &User{username, password}
}

func (u *User) Username() string {
	return u.username
}
func (u *User) Password() string {
	return u.password
}
func (u *User) SetUsername(username string) {
	u.username = username
}
func (u *User) SetPassword(password string) {
	u.password = password
}

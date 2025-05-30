package main

import "fmt"

type User struct {
	login    string
	password string
}

func NewUser(login string, password string) *User {
	return &User{login: login, password: password}
}

func (u *User) GetLogin() string {
	return u.login
}

func main() {
	user := NewUser("John", "123456")
	fmt.Println(user.GetLogin())
}

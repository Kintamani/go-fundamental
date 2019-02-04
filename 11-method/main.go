package main

import "fmt"

type users struct {
	ID       int
	name     string
	email    string
	password string
}

func (u *users) GetUser() (int, string, string, string) {
	return u.ID, u.name, u.email, u.password
}

func (u *users) GetNewUser(newName string) {
	u.name = newName
}

func main() {
	userDefault := users{
		ID:       1,
		name:     "JAKA",
		email:    "jaka@mail.com",
		password: "123123",
	}

	fmt.Println(userDefault.GetUser())

	// CHANGE NAME
	userDefault.GetNewUser("PRABAWA")
	fmt.Println(userDefault.GetUser())
}

package main

import (
	"fmt"
)

type User struct {
	email, password string
}

var (
	user, login []User
)

func create_acc(user_insert *User) {
	for {
		fmt.Println("Ketik 0 untuk keluar")
		fmt.Print("Username: ")
		fmt.Scan(&user_insert.email)

		if user_insert.email == "0" {
			break
		}

		fmt.Print("Password: ")
		fmt.Scan(&user_insert.password)

		n, state := len(user), false 


		for i := 0; i < n; i++ {
			if user_insert.password == user[i].password && user_insert.email == user[i].email {
				state = true
				break
			} 
		}

		if state == false {
			user = append(user, *user_insert)
			break
		}

		fmt.Println("Email dan Username telah terdaftar")

	}
}

func login_user(user_login User) bool {
	n := len(user)

	for i := 0; i < n; i++ {
		if user_login.password == user[i].password && user_login.email == user[i].email {
			login = append(login, user_login)
			return true
		}
	}
	return false
}
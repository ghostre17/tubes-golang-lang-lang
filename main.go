package main

import "fmt"

func main() {
	var (
		n    int
		s, l User
	)
	for {
		fmt.Println("===========================")
		fmt.Println("1. Create Acc")
		fmt.Println("2. Login")
		fmt.Println("3. cek login")
		fmt.Println("4. cek user terdaftar")
		fmt.Println("0. Exit")
		fmt.Println("===========================")
		fmt.Print("Masukkan kode: ")
		fmt.Scan(&n)
		switch n {
		case 1:
			create_acc(&s)
		case 2:
			for {
				fmt.Println("===========================")
				fmt.Println("Ketik 0 untuk keluar")
				fmt.Print("Masukkan username: ")
				fmt.Scan(&l.email)

				if l.email == "0" {
					break
				}

				fmt.Print("Masukkan password: ")
				fmt.Scan(&l.password)

				check := login_user(l)

				if check == true {
					dashboard()
					break
				}

				fmt.Println("Email dan password salah!")
			}

		case 3:
			fmt.Println(login)
		case 4:
			fmt.Println(user)
		case 0:
			return
		}

	}
}

func dashboard() {
	var (
		n int
		m Rating
	)

	for {
		fmt.Println("===========================")
		fmt.Println("1. Tambahkan film")
		fmt.Println("2. Update film")
		fmt.Println("3. Search film")
		fmt.Println("4. Explore film")
		fmt.Println("5. Delete film")
		fmt.Println("0. Exit")
		fmt.Println("===========================")
		fmt.Print("Masukkan kode: ")
		fmt.Scan(&n)
		switch n {
		case 1:
			create_rating(&m)
		case 2:
		case 3:
		case 4:
			sort_menu()
		case 5:
		case 0:
			login = nil
			return
		}

	}
}

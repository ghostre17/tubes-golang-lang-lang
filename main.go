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
		fmt.Println("2. Collections")
		fmt.Println("3. Search film")
		fmt.Println("4. Explore film")
		fmt.Println("5. Total Genre and Avg Rate")
		fmt.Println("0. Exit")
		fmt.Println("===========================")
		fmt.Print("Masukkan kode: ")
		fmt.Scan(&n)
		switch n {
		case 1:
			create_rating(&m)
		case 2:
			collections_output(rate)
		case 3:
			searching()
		case 4:
			sort_menu()
		case 5:
			avg_output()
		case 0:
			login = nil
			return
		}

	}
}

func sort_menu() {
	var n int8
	for {
		fmt.Println("\n========== DAFTAR FILM ==========")
		for i := 0; i < len(rate); i++ {
			fmt.Printf(
				"[%d]\nNama      : %s\nDeskripsi : %s\nGenre      : %s\nRilis     : %d\nRating    : %.1f\n\n",
				i+1,
				rate[i].movie,
				rate[i].description,
				rate[i].genre,
				rate[i].rilis,
				rate[i].rating,
			)
		}
		fmt.Println("=== Sorting By ===")
		fmt.Println("1. Rilis")
		fmt.Println("2. Rating")
		fmt.Println("0. Exit")
		fmt.Print("Masukkan kode: ")
		fmt.Scan(&n)
		switch n {
		case 1:
			var kode int
			for {
				fmt.Println("=== Sorting By Rilis ===")
				fmt.Println("1. Descending")
				fmt.Println("2. Ascending")
				fmt.Println("0. Exit")
				fmt.Scan(&kode)
				switch kode {
				case 1:
					fmt.Println("Descending")
				case 2:
					fmt.Println("Ascending")
				case 0:
					return
				}

				new := insertion(rate, kode)
				fmt.Println("===========================")
				for i := 0; i < len(new); i++ {
					fmt.Printf(
						"[%d]\nNama      : %s\nDeskripsi : %s\nGenre      : %s\nRilis     : %d\nRating    : %.1f\n\n",
						i+1,
						new[i].movie,
						new[i].description,
						new[i].genre,
						new[i].rilis,
						new[i].rating,
					)
				}
				fmt.Println("===========================")
			}
		case 2:
			var kode int
			for {
				fmt.Println("=== Sorting By Rating ===")
				fmt.Println("1. Descending")
				fmt.Println("2. Ascending")
				fmt.Println("0. Exit")
				fmt.Scan(&kode)
				switch kode {
				case 1:
					fmt.Println("=== Descending ===")
				case 2:
					fmt.Println("=== Ascending ===")
				case 0:
					fmt.Println("======")
					return
				}

				new := selection(rate, kode)
				fmt.Println("===========================")
				for i := 0; i < len(new); i++ {
					fmt.Printf(
						"[%d]\nNama      : %s\nDeskripsi : %s\nGenre      : %s\nRilis     : %d\nRating    : %.1f\n\n",
						i+1,
						new[i].movie,
						new[i].description,
						new[i].genre,
						new[i].rilis,
						new[i].rating,
					)
				}
				fmt.Println("===========================")
			}
		case 0:
			return
		}
	}
}

func collections_output(arr []Rating) {
	sorted := collections(arr)
	var n int

	for {
		fmt.Println("===========================")
		for i := 0; i < len(sorted); i++ {
			fmt.Printf(
				"[%d]\nNama      : %s\nDeskripsi : %s\nGenre     : %s\nRilis     : %d\nRating    : %.1f\n\n",
				i+1,
				sorted[i].movie,
				sorted[i].description,
				sorted[i].genre,
				sorted[i].rilis,
				sorted[i].rating,
			)
		}
		fmt.Println("===========================")

		fmt.Println("=== Collections Menu ===")
		fmt.Println("1. Update")
		fmt.Println("2. Delete")
		fmt.Println("0. Exit")
		fmt.Print("Masukkan kode: ")
		fmt.Scan(&n)
		switch n {
		case 1:
		case 2:
			var s string
			fmt.Print("Masukkan judul movie yang akan dihapus: ")
			fmt.Scan(&s)

			for i := 0; i < len(rate); i++ {
				if rate[i].movie == s {
					rate = append(rate[:i], rate[i+1:]...)
				}
			}
		case 0:
			return
		}
	}
}

func avg_output() {
	all_rate := average_rating(rate)
	group_by_genre := count_avgRating_perGenre(rate)
	sorted := bubble_avg(group_by_genre)
	fmt.Println("===========================")
	for i := 0; i < len(sorted); i++ {
		fmt.Printf(
			"[%d]\nGenre      : %s\nJumlah     : %d\nAvg Rating : %.2f\n\n",
			i+1,
			sorted[i].genre,
			sorted[i].count,
			sorted[i].rate,
		)
	}
	fmt.Println("===========================")
	fmt.Printf("Average Rating Semua Film : %.2f\n", all_rate)
}
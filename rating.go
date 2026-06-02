package main

import (
	"fmt"
)

type Rating struct {
	id_user            int
	movie, description string
	rilis              int
	rating             float64
}

var rate []Rating

func create_rating(film *Rating) {
	for {
		fmt.Println("===========================")
		fmt.Println("Ketik 0 untuk keluar")
		fmt.Print("Nama Movie: ")
		fmt.Scan(&film.movie)

		if film.movie == "0" {
			break
		}

		fmt.Print("Deskripsi: ")
		fmt.Scan(&film.description)

		fmt.Print("Rilis: ")
		fmt.Scan(&film.rilis)

		fmt.Print("rating: ")
		fmt.Scan(&film.rating)

		n, state := len(rate), false

		for i := 0; i < n; i++ {
			if film.movie == rate[i].movie {
				state = true
				break
			}
		}

		if state == false {
			for i := 0; i < len(user); i++ {
				if login[0].email == user[i].email {
					film.id_user = i
					rate = append(rate, *film)
					break
				}
			}

		}

		fmt.Println("Movie sudah terdaftar")

	}
}

// tahun rilis sorting
func insertion(arr []Rating, kode int) []Rating {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		if kode == 1 {
			// descending
			for j >= 0 && (arr[j].rilis < key.rilis ||
				arr[j].rilis == key.rilis && arr[j].movie > key.movie) {
				arr[j+1] = arr[j]
				j--
			}
		} else {
			// ascending
			for j >= 0 && (arr[j].rilis > key.rilis ||
				arr[j].rilis == key.rilis && arr[j].movie > key.movie) {
				arr[j+1] = arr[j]
				j--
			}
		}
		arr[j+1] = key
	}
	return arr
}

// rating sorting
func selection(arr []Rating, kode int) []Rating {
	for i := 0; i < len(arr); i++ {
		minIdx := i
		if kode == 1 {
			for j := i + 1; j < len(arr); j++ {
				if arr[j].rating > arr[minIdx].rating {
					minIdx = j
				}
			}
		} else {
			for j := i + 1; j < len(arr); j++ {
				if arr[j].rating < arr[minIdx].rating {
					minIdx = j
				}
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
	return arr
}

func sort_menu() {
	var n int8
	for {
		fmt.Println("\n========== DAFTAR FILM ==========")
		for i := 0; i < len(rate); i++ {
			fmt.Printf(
				"[%d]\nNama      : %s\nDeskripsi : %s\nRilis     : %d\nRating    : %.1f\n\n",
				i+1,
				rate[i].movie,
				rate[i].description,
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
						"[%d]\nNama      : %s\nDeskripsi : %s\nRilis     : %d\nRating    : %.1f\n\n",
						i+1,
						new[i].movie,
						new[i].description,
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
						"[%d]\nNama      : %s\nDeskripsi : %s\nRilis     : %d\nRating    : %.1f\n\n",
						i+1,
						new[i].movie,
						new[i].description,
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

package main

import (
	"fmt"
	"strings"
)

type Rating struct {
	id_user                   int
	movie, description, genre string
	rilis                     int
	rating                    float64
}

type genre_count_avg struct {
	genre string
	count int
	rate  float64
}

var rate = []Rating{
	{1, "Interstellar", "Perjalanan antariksa mencari planet baru.", "Sci-Fi", 2014, 4.9},
	{1, "Inception", "Pencurian informasi melalui dunia mimpi.", "Sci-Fi", 2010, 4.8},
	{2, "The Dark Knight", "Batman melawan Joker.", "Action", 2008, 5.0},
	{2, "Oppenheimer", "Kisah pencipta bom atom.", "Biography", 2023, 4.6},
	{3, "Parasite", "Keluarga miskin menyusup ke keluarga kaya.", "Thriller", 2019, 4.7},
	{3, "Avengers Endgame", "Pertempuran terakhir melawan Thanos.", "Action", 2019, 4.5},
	{4, "Titanic", "Kisah cinta di kapal Titanic.", "Romance", 1997, 4.2},
	{4, "La La Land", "Musisi dan aktris mengejar mimpi.", "Musical", 2016, 4.1},
	{5, "Spirited Away", "Petualangan di dunia roh.", "Animation", 2001, 4.9},
	{5, "Your Name", "Dua remaja bertukar tubuh.", "Animation", 2016, 4.6},
	{6, "The Conjuring", "Investigasi rumah berhantu.", "Horror", 2013, 4.0},
	{6, "Hereditary", "Teror dalam sebuah keluarga.", "Horror", 2018, 4.1},
	{7, "The Shawshank Redemption", "Harapan di balik jeruji.", "Drama", 1994, 5.0},
	{7, "Forrest Gump", "Perjalanan hidup Forrest.", "Drama", 1994, 4.8},
	{8, "Top Gun Maverick", "Pilot veteran kembali mengudara.", "Action", 2022, 4.4},
}

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

		fmt.Print("Genre: ")
		fmt.Scan(&film.genre)

		fmt.Print("Rilis: ")
		fmt.Scan(&film.rilis)

		for {
			fmt.Print("rating: ")
			fmt.Scan(&film.rating)

			if film.rating > 5 || film.rating < 0 {
				fmt.Println("Rentang rating hanya 0-5")
			} else {
				break
			}
		}

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
					fmt.Println("Movie berhasil ditambahkan!")
					break
				}
			}
		} else {
			fmt.Println("Movie sudah terdaftar")
			break
		}

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

// buat binary
func bubble(arr []Rating) []Rating {
	for i := 0; i < len(arr)-1; i++ {
		tukar := false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j].movie > arr[j+1].movie {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				tukar = true
			}
		}
		if !tukar {
			break
		}
	}
	return arr
}

// ============ Searchin =============
func sequential(arr []Rating) []Rating {
	var target string
	var hasil []Rating
	fmt.Print("Search by Judul/Genre: ")
	fmt.Scan(&target)

	for i := 0; i < len(arr); i++ {
		if strings.Contains(arr[i].movie, target) || strings.Contains(arr[i].genre, target) {
			hasil = append(hasil, arr[i])
		}
	}

	return hasil
}

func binary(arr []Rating) int {
	var target string
	ndukur, ndisor := len(arr)-1, 0

	fmt.Print("Search by Judul (Binary): ")
	fmt.Scan(&target)

	for ndisor <= ndukur {
		tengah := (ndisor + ndukur) / 2

		if arr[tengah].movie == target {
			return tengah
		} else if arr[tengah].movie < target {
			ndisor = tengah + 1
		} else {
			ndukur = tengah - 1
		}
	}
	return -1
}

func searching() {
	var n int
	fmt.Println("=== Searching By ===")
	fmt.Println("1. Sequential")
	fmt.Println("2. Binary")
	fmt.Println("0. Exit")
	fmt.Print("Masukkan kode: ")
	fmt.Scan(&n)
	switch n {
	case 1:
		search := sequential(rate)
		fmt.Println("===========================")
		for i := 0; i < len(search); i++ {
			fmt.Printf(
				"[%d]\nNama      : %s\nDeskripsi : %s\nGenre     : %s\nRilis     : %d\nRating    : %.1f\n\n",
				i+1,
				search[i].movie,
				search[i].description,
				search[i].genre,
				search[i].rilis,
				search[i].rating,
			)
		}
		fmt.Println("===========================")
	case 2:
		bubble := bubble(rate)
		search := binary(bubble)
		if search != -1 {
			fmt.Println("===========================")
			fmt.Printf("Nama      : %s\nDeskripsi : %s\nGenre     : %s\nRilis     : %d\nRating    : %.1f\n\n",
				bubble[search].movie,
				bubble[search].description,
				bubble[search].genre,
				bubble[search].rilis,
				bubble[search].rating,
			)
			fmt.Println("===========================")
		} else {
			fmt.Println("Movie tidak ada  :(")
		}
	case 0:
		return
	}
}

func average_rating(arr []Rating) float64 {
	var total float64
	for i := 0; i < len(arr); i++ {
		total += arr[i].rating
	}

	return total / float64(len(arr))
}

func count_avgRating_perGenre(arr []Rating) []genre_count_avg {
	var result []genre_count_avg

	for i := 0; i < len(arr); i++ {
		status := -1
		for j := 0; j < len(result); j++ {
			if result[j].genre == arr[i].genre {
				status = j
				break
			}
		}
		if status == -1 {
			result = append(result, genre_count_avg{genre: arr[i].genre, count: 1, rate: arr[i].rating})
		} else {
			result[status].count += 1
			result[status].rate += arr[i].rating
		}
	}

	for i := 0; i < len(result); i++ {
		result[i].rate = result[i].rate / float64(result[i].count)
	}

	return result
}

func bubble_avg(arr []genre_count_avg) []genre_count_avg {
	for i := 0; i < len(arr)-1; i++ {
		tukar := false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j].genre > arr[j+1].genre {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				tukar = true
			}
		}
		if !tukar {
			break
		}
	}
	return arr
}

func collections(arr []Rating) []Rating {

	idx := 0
	var coll []Rating

	for i := 0; i < len(user); i++ {
		if login[0].email == user[i].email && login[0].password == user[i].password {
			idx = i
			break
		}
	}

	for i := 0; i < len(arr); i++ {
		if arr[i].id_user == idx {
			coll = append(coll, arr[i])
		}
	}

	return coll
}



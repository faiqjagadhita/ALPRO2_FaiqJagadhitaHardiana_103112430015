package main

import "fmt"

func main() {
	var berat1, berat2, total float64

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&berat1, &berat2)

		if berat1 < 0 || berat2 < 0 || total > 150 {
			break
		}

		total += berat1 + berat2

		if berat1-berat2 >= 9 || berat2-berat1 >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
	}

	fmt.Println("Proses selesai.")
}

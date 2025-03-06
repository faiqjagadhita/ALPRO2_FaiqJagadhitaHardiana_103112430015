package main

import "fmt"

func main() {
	var berat int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&berat)

	kg := berat / 1000
	gram := berat % 1000

	hargakg := kg * 10000
	hargagram := 0

	if gram > 0 {
		if gram >= 500 {
			hargagram = gram * 5
		} else {
			hargagram = gram * 15
		}
	}

	if kg >= 10 {
		hargagram = 0
	}
	total := hargakg + hargagram

	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, gram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", hargakg, hargagram)
	fmt.Printf("Total biaya: Rp. %d\n", total)
}

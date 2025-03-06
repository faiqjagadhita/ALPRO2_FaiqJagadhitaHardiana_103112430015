package main

import "fmt"

func main() {
	var K int
	fmt.Print("Masukkan nilai K: ")
	fmt.Scan(&K)

	akar2 := 1.0
	for i := 0; i <= K; i++ {
		pembilang := (4*i + 2) * (4*i + 2)
		penyebut := (4*i + 1) * (4*i + 3)
		akar2 *= float64(pembilang) / float64(penyebut)
	}

	fmt.Printf("Nilai akar 2 = %.10f\n", akar2)
}
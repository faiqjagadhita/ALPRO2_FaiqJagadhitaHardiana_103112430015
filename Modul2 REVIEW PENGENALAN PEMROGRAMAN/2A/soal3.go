package main

import "fmt"

func main() {
	var jarijari float64
	fmt.Print("Jari-jari = ")
	fmt.Scan(&jarijari)

	volume := (4.0 / 3.0) * 3.1415926535 * (jarijari * jarijari * jarijari)
	luas := 4 * 3.1415926535 * (jarijari * jarijari)

	fmt.Printf("Bola dengan jari-jari %.0f memiliki volume %.4f dan luas kulit %.4f\n", jarijari, volume, luas)
}

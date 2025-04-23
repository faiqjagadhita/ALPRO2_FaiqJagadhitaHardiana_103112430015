package main

import "fmt"

const maxCalon = 20

func bacaSuara() []int {
	var suara int
	var suaraMasuk []int

	fmt.Println("Masukkan suara (akhiri dengan 0):")
	for {
		fmt.Scan(&suara)
		if suara == 0 {
			break
		}
		suaraMasuk = append(suaraMasuk, suara)
	}
	return suaraMasuk
}

func hitungSuaraSah(suaraMasuk []int) (int, []int) {
	var perolehan = make([]int, maxCalon+1)
	var totalSah int

	for _, s := range suaraMasuk {
		if s >= 1 && s <= maxCalon {
			perolehan[s]++
			totalSah++
		}
	}
	return totalSah, perolehan
}

func tampilkanHasil(totalMasuk, totalSah int, perolehan []int) {
	fmt.Printf("Suara masuk: %d\n", totalMasuk)
	fmt.Printf("Suara sah: %d\n", totalSah)

	for i := 1; i <= maxCalon; i++ {
		if perolehan[i] > 0 {
			fmt.Printf("%d: %d\n", i, perolehan[i])
		}
	}
}

func main() {
	suaraMasuk := bacaSuara()
	totalMasuk := len(suaraMasuk)
	totalSah, perolehan := hitungSuaraSah(suaraMasuk)

	tampilkanHasil(totalMasuk, totalSah, perolehan)
}

package main

import "fmt"

const maxCalon = 20

func ambilSuara() []int {
	var suara int
	var daftarSuara []int

	fmt.Println("Masukkan suara (akhiri dengan 0):")
	for {
		fmt.Scan(&suara)
		if suara == 0 {
			break
		}
		daftarSuara = append(daftarSuara, suara)
	}
	return daftarSuara
}

func hitungDanValidasi(daftarSuara []int) (int, []int) {
	var suaraValid = make([]int, maxCalon+1)
	var suaraSah int

	for _, suara := range daftarSuara {
		if suara >= 1 && suara <= maxCalon {
			suaraValid[suara]++
			suaraSah++
		}
	}
	return suaraSah, suaraValid
}

func cariPemenang(suaraValid []int) (int, int) {
	var ketua, wakil int
	var suaraTerbanyak, suaraKeduaTerbanyak int

	for i := 1; i <= maxCalon; i++ {
		suara := suaraValid[i]
		if suara > suaraTerbanyak {
			suaraKeduaTerbanyak = suaraTerbanyak
			wakil = ketua

			suaraTerbanyak = suara
			ketua = i
		} else if suara == suaraTerbanyak && i < ketua {
			suaraKeduaTerbanyak = suaraTerbanyak
			wakil = ketua

			ketua = i
		} else if suara > suaraKeduaTerbanyak && i != ketua {
			suaraKeduaTerbanyak = suara
			wakil = i
		} else if suara == suaraKeduaTerbanyak && i < wakil && i != ketua {
			wakil = i
		}
	}

	return ketua, wakil
}

func tampilkanPemenang(suaraMasuk, suaraSah int, ketua, wakil int) {
	fmt.Printf("Suara masuk: %d\n", suaraMasuk)
	fmt.Printf("Suara sah: %d\n", suaraSah)
	fmt.Printf("Ketua RT: %d\n", ketua)
	fmt.Printf("Wakil ketua: %d\n", wakil)
}

func main() {
	daftarSuara := ambilSuara()
	suaraMasuk := len(daftarSuara)

	suaraSah, suaraValid := hitungDanValidasi(daftarSuara)
	ketua, wakil := cariPemenang(suaraValid)

	tampilkanPemenang(suaraMasuk, suaraSah, ketua, wakil)
}

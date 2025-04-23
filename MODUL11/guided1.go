package main

import "fmt"

func cariBarang(daftar []string, x string) bool {
	for _, barang := range daftar {
		if barang == x {
			return true
		}
	}
	return false
}

func main () {

	daftarBarang := []string{"sabun", "sampo", "odol", "tisu", "minyak"}
	barangDicari := "odol"

	ditemukan := cariBarang(daftarBarang, barangDicari)

	fmt.Println(ditemukan)
}

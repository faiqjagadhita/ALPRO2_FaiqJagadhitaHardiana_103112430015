package main

import (
	"fmt"
)

type Mahasiswa struct {
	nama string
	nim  string
}

func binarySearch(mahasiswa []Mahasiswa, nim string) int {
	start := 0
	end := len(mahasiswa) - 1

	for start <= end {
		tengah := (start + end) / 2
		if mahasiswa[tengah].nim == nim {
			return tengah 
		} else if mahasiswa[tengah].nim < nim {
			start = tengah + 1 
		} else {
			end = tengah - 1 
		}
	}

	return -1 
}

func main() {
	
	mahasiswa := []Mahasiswa{
		{nama: "andi", nim: "220001"},
		{nama: "budi", nim: "220002"},
		{nama: "citra", nim: "220003"},
		{nama: "doni", nim: "220004"},
	}

	X := "220003"

	index := binarySearch(mahasiswa, X)

	if index != -1 {
		fmt.Printf("Indeks mahasiswa ditemukan: %d\n", index)
	} else {
		fmt.Println("Nim tidak ditemukan")
	}
}

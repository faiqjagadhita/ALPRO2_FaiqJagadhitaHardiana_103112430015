<h1 style="text-align: center;">Laporan Praktikum Modul 12&13<br></h1>

<p style="text-align: center;">Fa'iq Jagadhita Hadiana - 103112430015</p>

#### Soal 1

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Domino struct {
	sisiKiri  int
	sisiKanan int
	nilai     int
	isBalak   bool
}

type Dominoes struct {
	kartu   [28]Domino
	sisa    int
}

func buatSetDomino(d *Dominoes) {
	idx := 0
	for kiri := 0; kiri <= 6; kiri++ {
		for kanan := kiri; kanan <= 6; kanan++ {
			d.kartu[idx] = Domino{
				sisiKiri:  kiri,
				sisiKanan: kanan,
				nilai:     kiri + kanan,
				isBalak:   kiri == kanan,
			}
			idx++
		}
	}
	d.sisa = 28
}

func kocokKartu(d *Dominoes) {
	rand.Seed(time.Now().UnixNano())
	for i := d.sisa - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d.kartu[i], d.kartu[j] = d.kartu[j], d.kartu[i]
	}
}

func ambilKartu(d *Dominoes) Domino {
	if d.sisa == 0 {
		return Domino{-1, -1, -1, false}
	}
	d.sisa--
	return d.kartu[d.sisa]
}

func gambarKartu(k Domino, sisi int) int {
	if sisi == 1 {
		return k.sisiKiri
	} else if sisi == 2 {
		return k.sisiKanan
	}
	return -1
}

func nilaiKartu(k Domino) int {
	return k.nilai
}

func main() {
	var setDomino Dominoes

	buatSetDomino(&setDomino)
	kocokKartu(&setDomino)

	fmt.Println("Menampilkan 5 kartu domino pertama yang diambil:")
	for i := 0; i < 5; i++ {
		k := ambilKartu(&setDomino)
		fmt.Printf("Kartu %d: (%d,%d), Nilai: %d, Balak: %v\n", i+1, k.sisiKiri, k.sisiKanan, nilaiKartu(k), k.isBalak)
	}
}

```

![](output/1.png)

Penjelasan :
Program ini membuat sebuah mesin abstrak domino yang merepresentasikan satu set kartu domino dengan 28 kartu unik. Tipe data `Domino` menyimpan nilai sisi kiri dan kanan kartu, total nilai pip, serta apakah kartu tersebut balak (kedua sisi sama). Tipe data `Dominoes` adalah kumpulan kartu yang disimpan dalam array beserta jumlah kartu yang tersisa.

Prosedur `buatSetDomino` menginisialisasi kartu-kartu domino dari sisi 0 sampai 6 tanpa duplikasi (misalnya kartu 1-4 hanya satu, bukan dua). Fungsi `kocokKartu` melakukan pengacakan kartu menggunakan algoritma Fisher-Yates untuk menghasilkan urutan kartu yang acak setiap kali dijalankan. Fungsi `ambilKartu` mengeluarkan kartu terakhir dari set (seperti mengambil dari tumpukan) dan mengurangi jumlah kartu yang tersisa.

Fungsi `gambarKartu` mengembalikan nilai sisi tertentu dari sebuah kartu, sedangkan `nilaiKartu` mengembalikan total nilai pip dari kartu tersebut. Program utama (`main`) menginisialisasi set kartu, mengocoknya, dan mengambil lima kartu pertama untuk ditampilkan ke layar, lengkap dengan informasi sisi, nilai, dan status balak kartu.

Dengan struktur ini, program sudah memenuhi fungsi dasar mesin domino dan siap dikembangkan untuk operasi dan permainan yang lebih kompleks.

#### Soal 2

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type KartuDomino struct {
	kiri     int
	kanan    int
	total    int
	adalahBalak bool
}

type SetDomino struct {
	tumpukan [28]KartuDomino
	sisa     int
}

func buatDomino(s *SetDomino) {
	idx := 0
	for a := 0; a <= 6; a++ {
		for b := a; b <= 6; b++ {
			s.tumpukan[idx] = KartuDomino{
				kiri:        a,
				kanan:       b,
				total:       a + b,
				adalahBalak: a == b,
			}
			idx++
		}
	}
	s.sisa = 28
}

func acakTumpukan(s *SetDomino) {
	rand.Seed(time.Now().UnixNano())
	for i := s.sisa - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s.tumpukan[i], s.tumpukan[j] = s.tumpukan[j], s.tumpukan[i]
	}
}

func tarikKartu(s *SetDomino) KartuDomino {
	if s.sisa == 0 {
		return KartuDomino{-1, -1, -1, false}
	}
	s.sisa--
	return s.tumpukan[s.sisa]
}

func nilaiSisi(k KartuDomino, sisi int) int {
	if sisi == 1 {
		return k.kiri
	} else if sisi == 2 {
		return k.kanan
	}
	return -1
}

func hitungNilai(k KartuDomino) int {
	return k.total
}

func apakahCocok(k1, k2 KartuDomino) bool {
	return k1.kiri == k2.kiri ||
		k1.kiri == k2.kanan ||
		k1.kanan == k2.kiri ||
		k1.kanan == k2.kanan
}

func cariCocok(s *SetDomino, acuan KartuDomino) KartuDomino {
	fmt.Println("\nProses pencarian kartu yang cocok...")
	for s.sisa > 0 {
		k := tarikKartu(s)
		fmt.Printf("Ditarik: (%d,%d)\n", k.kiri, k.kanan)
		if apakahCocok(k, acuan) {
			return k
		}
	}
	return KartuDomino{-1, -1, -1, false}
}

func main() {
	var kumpulan SetDomino

	buatDomino(&kumpulan)
	acakTumpukan(&kumpulan)

	kartuAcuan := tarikKartu(&kumpulan)
	fmt.Printf("Kartu acuan: (%d,%d)\n", kartuAcuan.kiri, kartuAcuan.kanan)

	kartuHasil := cariCocok(&kumpulan, kartuAcuan)

	if kartuHasil.kiri == -1 {
		fmt.Println("Tidak ditemukan kartu yang cocok.")
	} else {
		fmt.Printf("\nKartu yang cocok: (%d,%d)\n", kartuHasil.kiri, kartuHasil.kanan)
	}
}

```

![](output/2.png)

Penjelasan :
Program Go ini mengimplementasikan operasi dasar untuk mengelola dan memproses satu set kartu domino menggunakan pendekatan mesin abstrak. Dimulai dengan mendefinisikan tipe data `KartuDomino` untuk merepresentasikan sebuah kartu yang terdiri dari dua sisi (kiri dan kanan), total nilai pip, serta apakah kartu tersebut merupakan balak (kedua sisi sama). Struktur `SetDomino` mewakili sekumpulan 28 kartu domino dan jumlah kartu yang tersisa.

Fungsi `buatDomino` digunakan untuk menginisialisasi semua kombinasi unik kartu domino tanpa duplikasi. Selanjutnya, `acakTumpukan` digunakan untuk mengacak urutan kartu dalam array menggunakan metode Fisher-Yates Shuffle dengan seed berbasis waktu saat ini. Fungsi `tarikKartu` digunakan untuk mengambil kartu satu per satu dari belakang tumpukan. `nilaiSisi` dan `hitungNilai` digunakan untuk mengakses sisi atau nilai total kartu. Fungsi `apakahCocok` membandingkan dua kartu untuk melihat apakah mereka memiliki sisi yang sama, dan `cariCocok` digunakan untuk mencari kartu pertama dari tumpukan yang cocok dengan kartu acuan.

Pada bagian `main`, program menjalankan semua proses tersebut mulai dari pembuatan set domino, pengacakan, penarikan kartu acuan, hingga pencarian kartu yang cocok. Program ini menggambarkan prinsip mesin abstrak dengan baik, di mana operasi kompleks dibangun dari komponen dan prosedur dasar yang saling mendukung.

#### Soal 3

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Domino struct {
	kiri  int
	kanan int
}

func buatSetKartuDomino() []Domino {
	var set []Domino
	i := 0
	for i <= 6 {
		j := i
		for j <= 6 {
			set = append(set, Domino{i, j})
			j++
		}
		i++
	}
	return set
}

func kocokKartu(kartu []Domino) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(kartu), func(i, j int) {
		kartu[i], kartu[j] = kartu[j], kartu[i]
	})
}

func cocok(kartu Domino, ujung int) bool {
	return kartu.kiri == ujung || kartu.kanan == ujung
}

func balikKartu(kartu Domino) Domino {
	return Domino{kartu.kanan, kartu.kiri}
}

func tampilkanKartu(kartu Domino) string {
	return fmt.Sprintf("[%d|%d]", kartu.kiri, kartu.kanan)
}

func main() {
	set := buatSetKartuDomino()
	ronde := 0
	skorTotal := 0

	for {
		ronde++
		fmt.Printf("\n--- Ronde %d ---\n", ronde)
		kocokKartu(set)

		tumpukan := append([]Domino(nil), set...)
		pemain := tumpukan[:7]
		tumpukan = tumpukan[7:]

		ranrai := []Domino{tumpukan[0]}
		tumpukan = tumpukan[1:]
		skor := 0

		for {
			fmt.Println("\nRantai:")
			i := 0
			for i < len(ranrai) {
				fmt.Print(tampilkanKartu(ranrai[i]), " ")
				i++
			}
			fmt.Println("\n\nKartu Anda:")
			i = 0
			for i < len(pemain) {
				fmt.Printf("%2d: %s\n", i+1, tampilkanKartu(pemain[i]))
				i++
			}
			fmt.Println("101010: Selesai Ronde | 909090: Selesai Permainan")
			fmt.Print("Pilihan Anda (-7..-1 kiri, 1..7 kanan): ")

			var pilih int
			fmt.Scan(&pilih)

			if pilih == 101010 {
				fmt.Println("Ronde dihentikan.")
				break
			} else if pilih == 909090 {
				fmt.Println("Permainan selesai.")
				fmt.Printf("Total Ronde: %d, Skor Total: %d\n", ronde-1, skorTotal)
				return
			}

			if pilih < -len(pemain) || pilih > len(pemain) || pilih == 0 {
				fmt.Println("Pilihan tidak valid.")
				continue
			}

			abs := pilih
			if abs < 0 {
				abs = -abs
			}
			if abs > len(pemain) {
				fmt.Println("Indeks kartu tidak tersedia.")
				continue
			}

			kartuPilihan := pemain[abs-1]
			bisaMain := false

			if pilih < 0 {
				ujung := ranrai[0].kiri
				if cocok(kartuPilihan, ujung) {
					if kartuPilihan.kanan == ujung {
						ranrai = append([]Domino{kartuPilihan}, ranrai...)
					} else {
						ranrai = append([]Domino{balikKartu(kartuPilihan)}, ranrai...)
					}
					bisaMain = true
				}
			} else {
				ujung := ranrai[len(ranrai)-1].kanan
				if cocok(kartuPilihan, ujung) {
					if kartuPilihan.kiri == ujung {
						ranrai = append(ranrai, kartuPilihan)
					} else {
						ranrai = append(ranrai, balikKartu(kartuPilihan))
					}
					bisaMain = true
				}
			}

			if bisaMain {
				pemain = append(pemain[:abs-1], pemain[abs:]...)
				skor++
			} else {
				if len(tumpukan) > 0 {
					fmt.Println("Kartu tidak cocok. Mengambil kartu dari tumpukan...")
					pemain = append(pemain, tumpukan[0])
					tumpukan = tumpukan[1:]
				} else {
					fmt.Println("Kartu tidak cocok dan tumpukan sudah habis.")
				}
			}

			if len(pemain) == 0 {
				fmt.Println("Semua kartu habis. Ronde selesai!")
				break
			}

			bisaMainLagi := false
			ujungKiri := ranrai[0].kiri
			ujungKanan := ranrai[len(ranrai)-1].kanan
			i = 0
			for i < len(pemain) {
				if cocok(pemain[i], ujungKiri) || cocok(pemain[i], ujungKanan) {
					bisaMainLagi = true
					break
				}
				i++
			}
			if !bisaMainLagi && len(tumpukan) == 0 {
				fmt.Println("Tidak ada kartu yang bisa dimainkan dan tumpukan habis. Ronde selesai!")
				break
			}
		}

		fmt.Printf("Skor Ronde %d: %d\n", ronde, skor)
		skorTotal += skor
	}
}
```

![](output/3.png)

Penjelasan :
Program Go ini mensimulasikan permainan **domino solitaire interaktif**. Pada setiap ronde, program membuat 28 kartu domino standar menggunakan fungsi `buatSetKartuDomino` dan mengacak urutannya dengan `kocokKartu`, menggunakan seed acak berbasis waktu. Pemain akan menerima tujuh kartu, sementara sisanya disisihkan sebagai tumpukan. Satu kartu dari tumpukan akan ditempatkan di meja sebagai permulaan rantai.

Pemain diminta untuk memilih kartu dari tangan mereka dan meletakkannya di ujung kiri (dengan input angka negatif) atau kanan (angka positif) dari rantai. Program memverifikasi kecocokan kartu menggunakan fungsi `cocok`. Bila diperlukan, kartu akan dibalik menggunakan `balikKartu` agar sisi yang cocok bisa ditempatkan secara valid.

Jika kartu dapat dipasang, maka skor ronde bertambah dan kartu tersebut dikeluarkan dari tangan pemain. Jika tidak cocok, dan masih ada kartu di tumpukan, pemain akan menarik satu kartu. Jika tumpukan kosong, pemain tidak bisa menambah kartu.

Ronde akan berakhir jika pemain memilih untuk menghentikan ronde (input `101010`), berhasil menghabiskan semua kartu di tangan, atau tidak ada lagi langkah yang dapat dilakukan dan tumpukan juga telah habis. Setelah itu, skor ronde ditambahkan ke skor total. Pemain dapat melanjutkan ke ronde berikutnya atau mengakhiri permainan sepenuhnya dengan memasukkan `909090`.

#### Soal 4

```go
package main

import (
	"fmt"
	"strings"
)

var input string
var indeks int
var karakterSekarang byte

// Menginisialisasi mesin karakter
func start() {
	indeks = 0
	karakterSekarang = input[indeks]
}

// Melanjutkan pembacaan ke karakter berikutnya
func maju() {
	indeks++
	if indeks < len(input) {
		karakterSekarang = input[indeks]
	}
}

// Mengambil karakter saat ini yang sedang dibaca
func cc() byte {
	return karakterSekarang
}

// Mengecek apakah sudah mencapai akhir untaian (ditandai titik)
func eop() bool {
	return karakterSekarang == '.'
}

func main() {
	fmt.Print("Masukkan teks berakhiran titik ('.'): ")
	fmt.Scanln(&input)

	if !strings.HasSuffix(input, ".") {
		fmt.Println("Kalimat harus diakhiri dengan tanda titik ('.')")
		return
	}

	start()

	totalKarakter := 0
	totalA := 0
	jumlahLE := 0
	var sebelumnya byte = 0

	for !eop() {
		karakter := cc()
		totalKarakter++

		if strings.ToUpper(string(karakter)) == "A" {
			totalA++
		}

		if strings.ToUpper(string(sebelumnya)) == "L" && strings.ToUpper(string(karakter)) == "E" {
			jumlahLE++
		}

		sebelumnya = karakter
		maju()
	}

	persentaseA := float64(totalA) / float64(totalKarakter) * 100

	fmt.Println("Total karakter terbaca :", totalKarakter)
	fmt.Println("Jumlah huruf 'A'       :", totalA)
	fmt.Printf("Persentase huruf 'A'   : %.2f%%\n", persentaseA)
	fmt.Println("Jumlah pasangan 'LE'   :", jumlahLE)
}

```

![](output/4.png)

Penjelasan :
Program di atas mengimplementasikan **mesin karakter sederhana** yang beroperasi pada sebuah string yang diakhiri dengan tanda titik (`.`). Program menyediakan **operasi dasar** seperti `start()` untuk mulai membaca dari awal, `maju()` untuk maju ke karakter berikutnya, `cc()` untuk mendapatkan karakter saat ini, dan `eop()` untuk mengecek akhir input. Program ini membaca setiap karakter dari input satu per satu dan menghitung jumlah total karakter (tidak termasuk titik), menghitung banyaknya huruf **'A'** (baik huruf kecil maupun besar), serta menghitung **kemunculan pasangan huruf 'L' diikuti 'E'** (kata "LE"). Setelah pembacaan selesai, program menghitung **persentase huruf 'A'** terhadap total karakter yang terbaca. Semua hasil kemudian ditampilkan ke layar. Program menjaga kesesuaian penuh dengan konsep mesin karakter seperti yang dijelaskan dalam soal.


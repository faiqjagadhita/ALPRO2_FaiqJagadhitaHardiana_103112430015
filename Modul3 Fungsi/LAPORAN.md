<h1 style="text-align: center;">Laporan Praktikum Modul 3<br>FUNGSI</h1>
<p style="text-align: center;">Fa'iq Jagadhita Hadiana - 103112430015</p>
___
### Soal 1

```go
package main

import "fmt"
  
func factorial(n int) int {
    if n == 0 || n == 1 {
        return 1
    }
    hasil := 1
    for i := 2; i <= n; i++ {
        hasil *= i
    }
    return hasil
}
  
func permutation(n, r int) int {
    if n < r {
        return 0
    }
    hasil := factorial(n) / factorial(n-r)
    return hasil
}
  
func combination(n, r int) int {
    if n < r {
        return 0
    }
    hasil := factorial(n) / (factorial(r) * factorial(n-r))
    return hasil
}
 
func main() {
    var a, b, c, d int
    fmt.Scan(&a, &b, &c, &d)
    
    if a >= c && b >= d {
        fmt.Println(permutation(a, c), combination(a, c))
        fmt.Println(permutation(b, d), combination(b, d))
    }
}
```
![](1.png)

Penjelasan :
Program Go di atas menghitung permutasi dan kombinasi dari dua pasang bilangan yang dimasukkan oleh pengguna. Fungsi factorial menghitung nilai faktorial dari bilangan bulat positif menggunakan perulangan. Fungsi permutation menghitung jumlah susunan berbeda dari r elemen yang diambil dari n elemen dengan rumus P(n, r) = n! / (n - r)!, sedangkan fungsi combination menghitung jumlah cara memilih r elemen dari n elemen tanpa memperhatikan urutan dengan rumus C(n, r) = n! / (r! × (n - r)!). Pada fungsi main, program membaca empat bilangan bulat a, b, c, dan d dari input pengguna, lalu memeriksa apakah a ≥ c dan b ≥ d. Jika syarat tersebut terpenuhi, program menampilkan hasil perhitungan permutasi dan kombinasi untuk pasangan (a, c) dan (b, d) secara terpisah.

___
### Soal 2

```go
package main
  
import "fmt"

func f(x int) int {
    hasil := x * x
    return hasil
}
  
func g(x int) int {
    hasil := x - 2
    return hasil
}
  
func h(x int) int {
    hasil := x + 1
    return hasil
}

func main() {
    var a, b, c int
    fmt.Scan(&a, &b, &c)
    
    fogoh := f(g(h(a)))
    gohof := g(h(f(b)))
    hofog := h(f(g(c)))
  
    fmt.Print(fogoh, "\n", gohof, "\n", hofog, "\n")
}
```
![](2.png)

Penjelasan : 
Program Go di atas menghitung hasil komposisi dari tiga fungsi, yaitu f(x), g(x), dan h(x), yang masing-masing memiliki operasi sederhana. Fungsi f(x) menghitung kuadrat dari x, g(x) mengurangkan 2 dari x, dan h(x) menambahkan 1 ke x. Pada fungsi main, program membaca tiga bilangan bulat a, b, dan c dari input pengguna, lalu menghitung tiga komposisi fungsi: f(g(h(a))), g(h(f(b))), dan h(f(g(c))). Hasil dari masing-masing komposisi dicetak dalam baris terpisah.

___
### Soal 3

```go
package main

import (
    "fmt"
    "math"
)
 
func jarak(a, b, c, d float64) float64 {
    hasil := math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
    return hasil
}
  
func didalam(cx, cy, r, x, y float64) bool {
    hasil := jarak(cx, cy, x, y) <= r
    return hasil
} 

func main() {
    var cx1, cy1, r1, cx2, cy2, r2, x, y float64
    var didalam1, didalam2 bool

    fmt.Scan(&cx1, &cy1, &r1)
    fmt.Scan(&cx2, &cy2, &r2)
    fmt.Scan(&x, &y)

    didalam1 = didalam(cx1, cy1, r1, x, y)
    didalam2 = didalam(cx2, cy2, r2, x, y)

    if didalam1 && didalam2 {
        fmt.Println("Titik di dalam lingkaran 1 dan 2")
    } else if didalam1 {
        fmt.Println("Titik di dalam lingkaran 1")
    } else if didalam2 {
        fmt.Println("Titik di dalam lingkaran 2")
    } else {
        fmt.Println("Titik di luar lingkaran 1 dan 2")
    }
}
```
![](3.png)

Penjelasan :
Program Go di atas buat ngecek posisi sebuah titik terhadap dua lingkaran. Fungsi **jarak** dipakai buat ngitung jarak antara dua titik di bidang kartesian, sedangkan fungsi **didalam** ngecek apakah titik (x, y) ada di dalam atau pas di tepi lingkaran dengan pusat (cx, cy) dan jari-jari r. Di fungsi **main**, program bakal baca input untuk pusat dan jari-jari dua lingkaran, plus koordinat titik yang mau dicek. Setelah itu, dicek deh apakah titiknya ada di dalam lingkaran pertama, kedua, atau malah di luar keduanya, terus hasilnya ditampilin sesuai kondisi.
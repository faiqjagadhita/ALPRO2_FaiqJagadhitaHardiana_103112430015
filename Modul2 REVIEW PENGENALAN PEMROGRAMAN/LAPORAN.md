<h1 style="text-align: center;">Laporan Praktikum Modul 2<br>REVIEW PENGENALAN PEMROGRAMAN</h1>
<p style="text-align: center;">Fa'iq Jagadhita Hadiana - 103112430015</p>
___
### Soal 2A

1. Program Pertama
```go
package main

  

import "fmt"

  

func main() {

    var (

        satu, dua, tiga string

        temp string

    )

    fmt.Print("Masukan input string: ")

    fmt.Scanln(&satu)

    fmt.Print("Masukan input string: ")

    fmt.Scanln(&dua)

    fmt.Print("Masukan input string: ")

    fmt.Scanln(&tiga)

    fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)

    temp = satu

    satu = dua

    dua = tiga

    tiga = temp

    fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)

}
```

![](output/2a1.png)

Penjelasan : Setelah saya telusuri, program di atas meminta pengguna untuk memasukkan tiga string secara berurutan. Setelah semua string dimasukkan, program menampilkan urutan awal dari ketiga string tersebut. Selanjutnya, program melakukan pertukaran posisi string: string pertama dipindahkan ke posisi kedua, string kedua berpindah ke posisi ketiga, dan string ketiga ditempatkan di posisi pertama. Untuk membantu dalam proses pertukaran ini, program menggunakan variabel sementara `temp`. Setelah pertukaran selesai, program mencetak hasil akhir dengan urutan string yang sudah berubah.

2. Tahun kabisat
```go
package main

  

import "fmt"

  

func main() {

    var tahun int

    fmt.Print("Masukkan Tahun: ")

    fmt.Scan(&tahun)

  

    if tahun%400 == 0 || (tahun%4 == 0 && tahun%100 != 0) {

        fmt.Println("Tahun Kabisat: true")

    } else {

        fmt.Println("Tahun Kabisat: false")

    }

}
```

![](output/2a2.png)

Penjelasan : Program ini berfungsi untuk menentukan apakah suatu tahun merupakan tahun kabisat atau tidak. Pengguna diminta memasukkan sebuah tahun, lalu program memeriksa apakah tahun tersebut memenuhi aturan tahun kabisat dalam kalender Gregorian. Tahun dikatakan kabisat jika habis dibagi 400 atau jika habis dibagi 4 tetapi tidak habis dibagi 100. Jika salah satu dari kondisi tersebut terpenuhi, program mencetak "Tahun Kabisat: true", sedangkan jika tidak, program mencetak "Tahun Kabisat: false".

3. Program Bola
```go
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
```

![](output/2a3.png)

Penjelasan : Program ini menghitung volume dan luas permukaan sebuah bola berdasarkan jari-jari yang dimasukkan oleh pengguna. Setelah pengguna memasukkan nilai jari-jari, program menggunakan rumus volume bola yaitu (4/3) * π * r³ dan luas permukaan bola yaitu 4 * π * r² dengan π (pi) mendekati 3.1415926535. Hasil perhitungan kemudian ditampilkan dalam format desimal, di mana jari-jari ditampilkan sebagai bilangan bulat, sementara volume dan luas permukaan ditampilkan dengan empat angka di belakang koma.

4. suhu
```go
package main

  

import "fmt"

  

func main() {

    var celsius float64

    fmt.Print("Temperatur Celsius: ")

    fmt.Scan(&celsius)

  

    reamur := (4.0 / 5.0) * celsius

    fahrenheit := (9.0 / 5.0) * celsius + 32

    kelvin := celsius + 273

  

    fmt.Printf("Derajat Reamur: %.0f\n", reamur)

    fmt.Printf("Derajat Fahrenheit: %.0f\n", fahrenheit)

    fmt.Printf("Derajat Kelvin: %.0f\n", kelvin)

}
```

![[2a4.png]]

Penjelasan : Program ini mengonversi suhu dari skala Celsius ke tiga skala suhu lainnya: Reamur, Fahrenheit, dan Kelvin. Pengguna diminta memasukkan suhu dalam derajat Celsius, lalu program menghitung konversinya menggunakan rumus: Reamur = (4/5) × Celsius, Fahrenheit = (9/5) × Celsius + 32, dan Kelvin = Celsius + 273.

5. Program ASCII
```go
package main

  

import (

    "fmt"

)

  

func main() {

    var a, b, c, d, e int

    var x, y, z byte

  

    fmt.Scan(&a, &b, &c, &d, &e)

    fmt.Scanln()

  

    fmt.Scanf("%c%c%c", &x, &y, &z)

  

    fmt.Printf("%c%c%c%c%c\n", a, b, c, d, e)

    fmt.Printf("%c%c%c\n", x+1, y+1, z+1)

}
```

![](output/2a5.png)

Penjelasan : Program ini menerima lima angka dan tiga karakter dari input pengguna. Lima angka pertama diperlakukan sebagai kode ASCII, kemudian dikonversi menjadi karakter dan ditampilkan sebagai sebuah teks. Setelah itu, program membaca tiga karakter tambahan tanpa spasi, lalu menaikkan setiap karakter satu tingkat dalam kode ASCII sebelum mencetak hasil akhirnya. Jadi, jika pengguna memasukkan angka yang mewakili huruf dalam ASCII serta karakter seperti 'A', 'B', 'C', maka outputnya akan menampilkan teks sesuai angka yang dimasukkan dan karakter yang telah bergeser satu tingkat ke depan dalam ASCII.
___
### Soal 2B

1. Warna
```go
package main

  

import "fmt"

  

func main() {

    var warna1, warna2, warna3, warna4 string

    berhasil := true

  

    for i := 1; i <= 5; i++ {

        fmt.Printf("Percobaan %d: ", i)

        fmt.Scan(&warna1, &warna2, &warna3, &warna4)

  

        if warna1 != "merah" || warna2 != "kuning" || warna3 != "hijau" || warna4 != "ungu" {

            berhasil = false

        }

    }

  

    fmt.Println("BERHASIL:", berhasil)

}
```

![](output/2b1.png)
 
 Penjelasan : Program ini meminta pengguna untuk memasukkan empat warna sebanyak lima kali dalam sebuah perulangan. Pada setiap percobaan, program membaca empat warna yang diinput oleh pengguna dan membandingkannya dengan urutan yang sudah ditentukan, yaitu "merah", "kuning", "hijau", dan "ungu". Jika dalam salah satu percobaan ada warna yang tidak sesuai, variabel berhasil diubah menjadi false. Setelah lima percobaan selesai, program mencetak hasil akhir dengan menampilkan "BERHASIL: true" jika semua input sesuai di setiap percobaan, atau "BERHASIL: false" jika ada satu saja percobaan yang tidak memenuhi urutan warna yang diharapkan.

2. 
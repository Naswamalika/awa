package main

import (
	"fmt"
)

// cetakDeret mencetak setiap suku dari deret yang dimulai dari nilai awal n
func cetakDeret(n int) {
	for n != 1 {
		fmt.Printf("%d ", n) // Cetak suku saat ini diikuti oleh spasi
		if n%2 == 0 {
			n /= 2 // Jika genap, bagi dua
		} else {
			n = 3*n + 1 // Jika ganjil, kalikan tiga lalu tambah satu
		}
	}
	fmt.Println(1) // Cetak 1 sebagai akhir deret dan memberhentikan rekursi
}

func deretBilangan() {
	// Input: Suku awal (bilangan positif kurang dari 1000000)
	var n int
	fmt.Println("Latihan 3 Prosedur - Deret Bilangan")
	fmt.Println("Masukkan nilai suku awal (bilangan positif < 1000000):")
	fmt.Scan(&n)

	// Validasi input
	if n <= 0 || n >= 1000000 {
		fmt.Println("Masukan tidak valid, masukkan bilangan positif kurang dari 1000000.")
		return
	}

	// Panggil prosedur cetakDeret
	cetakDeret(n)
}
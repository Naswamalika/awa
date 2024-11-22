package main

import (
	"fmt"
)

func sqrt2(k int) float64 {
	// Mulai dengan nilai produk awal = 1.0
	product := 1.0

	// Iterasi dari 0 hingga K
	for i := 0; i <= k; i++ {
		numerator := float64((4*i + 2) * (4*i + 2))   // (4k + 2)^2, := berfungsi untuk mendeklarasikan variabel numerator dengan tipe data float64
		denominator := float64((4*i + 1) * (4*i + 3)) // (4k + 1)(4k + 3)
		product *= numerator / denominator // *= berfungsi untuk mengalikan nilai product dengan nilai numerator/denominator
	}

	return product
}

func perulanganLatihan4() {
	var k int

	// Input nilai K dari pengguna
	fmt.Print("Masukkan nilai K: ")
	fmt.Scan(&k)

	// Hitung hampiran sqrt(2)
	result := sqrt2(k)

	// Tampilkan hasil dengan presisi 10 angka di belakang koma
	fmt.Printf("Nilai hampiran akar 2 = %.10f\n", result)
}
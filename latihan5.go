package main

import (
	"fmt"
)

func hitungBiaya(berat int) int {
	// Konversi berat dalam gram menjadi kg dan sisa gram
	kg := berat / 1000 // berat dibagi 1000 untuk mendapatkan kg
	sisa := berat % 1000 // % adalah operator modulus, sisa diambil dari pembagian berat dengan 1000

	// Biaya dasar berdasarkan jumlah kg, kg dikali 10000 untuk mendapatkan biaya
	biaya := kg * 10000

	// Jika berat lebih dari 10 kg, sisa gram tidak dikenakan biaya
	if kg >= 10 {
		return biaya
	}

	// Jika sisa >= 500 gram, tambahan biaya Rp. 5,- per gram
	if sisa >= 500 {
		biaya += sisa * 5
	} else {
		// Jika sisa < 500 gram, tambahan biaya Rp. 15,- per gram
		biaya += sisa * 15
	}

	return biaya
}

func percabanganLatihan1() {
	var berat int

	// Input berat parsel dalam gram
	fmt.Print("Masukkan berat parsel (dalam gram): ")
	fmt.Scan(&berat)

	// Hitung biaya kirim
	totalBiaya := hitungBiaya(berat)

	// Hitung jumlah kg dan sisa gram
	kg := berat / 1000
	sisa := berat % 1000

	// Tampilkan hasil perhitungan
	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, sisa)
	fmt.Printf("Detail biaya: Rp.%d + Rp.%d\n", kg*10000, totalBiaya-(kg*10000))
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
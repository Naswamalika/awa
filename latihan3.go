package main

import (
	"fmt"
	"math"
)

func perulanganLatihan3() {
	var weight1, weight2 float64

	for {
		// Masukkan berat belanjaan
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&weight1, &weight2)

		// Periksa apakah ada berat yang negatif, jika ada maka hentikan program
		if weight1 < 0 || weight2 < 0 {
			break
		}

		// Hitung perbedaan antara kedua berat, jika lebih dari 9 maka motor akan oleng
		diff := math.Abs(weight1 - weight2)

		// Periksa apakah motor akan oleng
		if diff >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}

		// Periksa apakah total berat melebihi 150 kg, jika ya maka hentikan program
		if weight1+weight2 > 150 {
			break
		}
	}

	fmt.Println("Proses selesai.")
}
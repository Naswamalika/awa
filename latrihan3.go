package main

import "fmt"

type arrBalita [100]float64

// Subprogram untuk menghitung berat minimum dan maksimum
func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	// Inisialisasi nilai minimum dan maksimum dengan elemen pertama array
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	// Loop untuk mencari nilai minimum dan maksimum
	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

// Fungsi rerata untuk menghitung rata-rata berat balita
func rerata(arrBerat arrBalita, n int) float64 {
	total := 0.0 // Inisialisasi total berat dengan 0
	for i := 0; i < n; i++ { // Loop untuk menjumlahkan total berat
		total += arrBerat[i] // Menambahkan setiap elemen array ke total
	}
	return total / float64(n)
}

func latihan3() {
	var n int
	var berat arrBalita
	var bMin, bMax float64

	fmt.Println("-=-=-=- JORDAN'S LATIHAN 3 - BERAT BALITA PROGRAM -=-=-=-")
	fmt.Print("Welcome! Pertama, anda punya berapa balita yang ingin diukur beratnya? ")
	fmt.Scan(&n)

	// Perulangan for untuk meminta input berat balita sebanyak n kali
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	// Memanggil fungsi hitungMinMax untuk mendapatkan berat minimum dan maksimum dan rerata untuk mendapatkan rata-rata berat balita
	hitungMinMax(berat, n, &bMin, &bMax)
	rataRata := rerata(berat, n)

	fmt.Println("\n-=-=-=- HASIL PERHITUNGAN BERAT BALITA -=-=-=-")
	fmt.Printf("Berat balita minimum : %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum : %.2f kg\n", bMax)
	fmt.Printf("Rerata berat balita : %.2f kg\n", rataRata)
}
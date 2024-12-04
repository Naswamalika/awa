package main

import "fmt"

func latihan1() {
	var N int
	fmt.Println("-=-=-=- JORDAN'S LATIHAN 1 - RABBIT WEIGHT PROGRAM -=-=-=-")
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&N)

	if N <= 0 || N > 1000 {
		fmt.Println("Jumlah anak kelinci harus antara 1 dan 1000")
		return
	}

	weights := make([]float64, N)
	fmt.Println("Masukkan berat anak kelinci:")
	for i := 0; i < N; i++ {
		fmt.Scan(&weights[i])
	}

	minWeight, maxWeight := weights[0], weights[0]

	for _, weight := range weights {
		if weight < minWeight {
			minWeight = weight
		}
		if weight > maxWeight {
			maxWeight = weight
		}
	}

	fmt.Println("\n-=-=-=- HASIL PERHITUNGAN -=-=-=-")
	fmt.Printf("Berat kelinci terkecil: %.2f\n", minWeight)
	fmt.Printf("Berat kelinci terbesar: %.2f\n", maxWeight)
}
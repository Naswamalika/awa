package main

import "fmt"

// Kapasitas maksimum array berat ikan
const maxCapacity = 1000

// Fungsi untuk menghitung total berat ikan di setiap wadah dan rata-rata berat per wadah
func calculateWeight(x int, y int, weights [maxCapacity]float64) ([]float64, float64) {
	// Hitung jumlah wadah yang diperlukan
	numContainers := (x + y - 1) / y
	containerWeights := make([]float64, numContainers)

	// Mengisi setiap wadah dengan total berat ikan sesuai kapasitas y
	for i := 0; i < x; i++ {
		containerIndex := i / y
		containerWeights[containerIndex] += weights[i]
	}

	// Menghitung rata-rata berat ikan per wadah
	totalWeight := 0.0
	for _, weight := range containerWeights {
		totalWeight += weight
	}
	averageWeight := totalWeight / float64(numContainers)

	return containerWeights, averageWeight
}

func latihan2() {
	var x, y int
	var weights [maxCapacity]float64

	fmt.Println("-=-=-=- JORDAN'S LATIHAN 2 - FISH WEIGHT PROGRAM -=-=-=-")
	fmt.Print("Hello! Pertama, masukkan jumlah ikan (x) dan kapasitas wadah (y): ")
	fmt.Scan(&x, &y)

	// Validasi input agar x tidak melebihi kapasitas maksimum
	if x > maxCapacity {
		fmt.Printf("Huh?? Banyak banget? Jumlah ikan melebihi kapasitas maksimum %d\n", maxCapacity)
		return
	}

	// Input berat tiap ikan, perulangan for untuk meminta input sebanyak x kali
	fmt.Printf("Ok nice! Lalu, masukkan berat ikan sebanyak %d:\n", x)
	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&weights[i])
	}

	// Memanggil fungsi untuk menghitung total berat ikan per wadah dan rata-rata berat
	containerWeights, averageWeight := calculateWeight(x, y, weights)

	// Perulangan for untuk menampilkan total berat ikan di setiap wadah
	fmt.Println("\n-=-=-=- HASIL PERHITUNGAN IKAN -=-=-=-")
	fmt.Println("Sip! Total berat ikan di setiap wadah adalah sebagai berikut:")
	for i, weight := range containerWeights {
		fmt.Printf("Wadah %d: %.2f kg\n", i+1, weight)
	}

	fmt.Printf("\nBerat rata-rata ikan di setiap wadah adalah %.2f kg\n", averageWeight)
}
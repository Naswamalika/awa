package main

import (
	"fmt"
	"math"
)

func latihan2() {
	var n int
	fmt.Println("-=-=- PENGISIAN ARRAY INTEGER -=-=-")
	fmt.Print("Masukkan jumlah elemen array (N): ")
	fmt.Scan(&n)

	// Membuat array dengan kapasitas n elemen
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i)
		fmt.Scan(&arr[i])
	}

	// a. Menampilkan keseluruhan isi dari array
	fmt.Println("\n-=-=- HASIL OUTPUT -=-=-")
	fmt.Println("Isi keseluruhan array:", arr)

	// b. Menampilkan elemen-elemen array dengan indeks ganjil saja
	fmt.Print("Elemen dengan indeks ganjil: ")
	for i := 1; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// c. Menampilkan elemen-elemen array dengan indeks genap saja
	fmt.Print("Elemen dengan indeks genap: ")
	for i := 0; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// d. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x
	var x int
	fmt.Print("\n")
	fmt.Println("-=-=- KELIPATAN INDEKS -=-=-")
	fmt.Print("Masukkan nilai x untuk indeks kelipatan: ")
	fmt.Scan(&x)
	fmt.Println("\n-=-=- HASIL OUTPUT -=-=-")
	fmt.Printf("Elemen dengan indeks kelipatan %d: ", x)
	for i := x; i < len(arr); i += x {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// e. Menghapus elemen array pada indeks tertentu
	var delIdx int
	fmt.Print("\n")
	fmt.Println("-=-=- PENGHAPUSAN ELEMEN ARRAY -=-=-")
	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&delIdx)
	fmt.Println("\n-=-=- HASIL OUTPUT -=-=-")
	if delIdx >= 0 && delIdx < len(arr) {
		arr = append(arr[:delIdx], arr[delIdx+1:]...) // Menghapus elemen array pada indeks tertentu
		fmt.Println("Isi array setelah penghapusan:", arr)
	} else {
		fmt.Println("Indeks tidak valid, where!?")
	}

	// f. Menampilkan rata-rata dari bilangan yang ada di dalam array
	var sum int
	for _, val := range arr {
		sum += val
	}
	average := float64(sum) / float64(len(arr)) // len(arr) adalah panjang array
	fmt.Printf("Rata-rata elemen array: %.2f\n", average)

	// g. Menampilkan standar deviasi atau simpangan baku dari bilangan yang ada di dalam array
	var varianceSum float64
	for _, val := range arr { // Menghitung varians
		varianceSum += math.Pow(float64(val)-average, 2) // (x - rata-rata)^2
	}
	standardDeviation := math.Sqrt(varianceSum / float64(len(arr))) // Standar deviasi = akar dari varians
	fmt.Printf("Standar deviasi elemen array: %.2f\n", standardDeviation)

	// h. Menampilkan frekuensi dari suatu bilangan tertentu
	var target int
	fmt.Print("\n")
	fmt.Println("-=-=- FREKUENSI KEMUNCULAN BILANGAN -=-=-")
	fmt.Print("Masukkan bilangan untuk mencari frekuensinya: ")
	fmt.Scan(&target)
	frequency := 0 // Inisialisasi frekuensi adalah 0
	for _, val := range arr { // untuk setiap elemen array
		if val == target { // jika elemen array sama dengan target
			frequency++ // tambahkan frekuensi
		}
	}
	fmt.Println("\n-=-=- HASIL OUTPUT -=-=-")
	fmt.Printf("Frekuensi kemunculan %d: %d kali\n", target, frequency)
	fmt.Print("\n")
}
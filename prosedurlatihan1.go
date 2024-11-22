package main

import (
	"fmt"
)

// Fungsi untuk menghitung faktorial
func prosedurFactorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// Fungsi untuk menghitung permutasi
func prosedurPermutation(n, r int) int {
	return prosedurFactorial(n) / prosedurFactorial(n-r)
}

// Fungsi untuk menghitung kombinasi
func prosedurCombination(n, r int) int {
	return prosedurFactorial(n) / (prosedurFactorial(r) * prosedurFactorial(n-r))
}

func perDanKomProsedur() {
	var a, b, c, d int

	fmt.Print("Latihan 1 Prosedur - Permutasi dan Kombinasi\n")
	// Meminta input dari pengguna
	fmt.Println("Masukkan nilai a, b, c, d: ")
	fmt.Scan(&a, &b, &c, &d)

	// Menghitung permutasi dan kombinasi untuk a terhadap c
	p1 := prosedurPermutation(a, c)
	c1 := prosedurCombination(a, c)

	// Menghitung permutasi dan kombinasi untuk b terhadap d
	p2 := prosedurPermutation(b, d)
	c2 := prosedurCombination(b, d)

	// Output hasil
	fmt.Printf("P(%d,%d) = %d\n", a, c, p1)
	fmt.Printf("C(%d,%d) = %d\n", a, c, c1)
	fmt.Printf("P(%d,%d) = %d\n", b, d, p2)
	fmt.Printf("C(%d,%d) = %d\n", b, d, c2)
}
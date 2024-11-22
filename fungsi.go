package main

import (
	"fmt"
)

func fungsi() {
	var a, b int

	fmt.Print("Contoh Fungsi - Permutasi\n")
	fmt.Print("Masukkan nilai a dan b: ")
	fmt.Scan(&a, &b)
	if a >= b {
		fmt.Println(permutasi(a, b))
	} else {
		fmt.Println(permutasi(b, a))
	}
}

func factorialfungsi(n int) int {
	var hasil int = 1
	var i int
	for i = 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}

func permutasi(n, r int) int {
	return factorialfungsi(n) / factorialfungsi(n-r)
}
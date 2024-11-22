package main

import (
    "fmt"
)

// Mencari faktor-faktor dari bilangan n
func findFactors(n int) []int {
    var factors []int
    for i := 1; i <= n; i++ { // Mencari faktor dari 1 hingga n
        if n%i == 0 {
            factors = append(factors, i) // Menambahkan faktor ke slice, slice maksudnya adalah array yang bisa berubah-ubah panjangnya
        }
    }
    return factors
}

// Menentukan apakah bilangan n prima
func isPrime(n int) bool {
    if n < 2 {
        return false
    }
    for i := 2; i*i <= n; i++ { // Mencari faktor dari 2 hingga akar n
        if n%i == 0 {
            return false
        }
    }
    return true
}

// Fungsi utama
func percabanganLatihan3() {
    var b int
    fmt.Print("Bilangan: ")
    fmt.Scan(&b)

    if b <= 0 {
        fmt.Println("Bilangan harus lebih dari 0")
        return
    }

    // Mencari faktor-faktor dari b
    fmt.Print("Faktor: ")
    factors := findFactors(b)
    for _, factor := range factors {
        fmt.Print(factor, " ")
    }
    fmt.Println()

    // Menentukan apakah bilangan prima
    isPrime := isPrime(b)
    fmt.Println("Prima:", isPrime)
}
package main

import (
    "fmt"
)

// Fungsi dasar
func f(x int) int {
    return x * x // fungsi ini mengembalikan nilai x pangkat 2
}

func g(x int) int {
    return x - 2 // fungsi ini mengembalikan nilai x dikurangi 2
}

func h(x int) int {
    return x + 1 // fungsi ini mengembalikan nilai x ditambah 1
}

// Fungsi komposisi sesuai instruksi modul
func fogoh(x int) int {
    return f(g(h(x))) // Alasan fogoh bisa menjadi f(g(h(x))) adalah karena h(x), g(x), dan f(x) adalah fungsi yang bisa dijalankan secara berurutan
}

func gohof(x int) int {
    return g(h(f(x))) // Alasan gohof bisa menjadi g(h(f(x))) adalah karena f(x), h(x), dan g(x) adalah fungsi yang bisa dijalankan secara berurutan
}

func hofog(x int) int {
    return h(f(g(x))) // Alasan hofog bisa menjadi h(f(g(x))) adalah karena g(x), f(x), dan h(x) adalah fungsi yang bisa dijalankan secara berurutan
}

func matematika() {
    var a, b, c int

    fmt.Print("Latihan 2 Fungsi - Fungsi Matematika\n")

    fmt.Print("Masukkan nilai a, b, c: ")
    fmt.Scanf("%d %d %d", &a, &b, &c)

    fmt.Println("fogoh(",a,") =", fogoh(a))  // fogoh(a)
    fmt.Println("gohof(",b,") =", gohof(b))  // gohof(b)
    fmt.Println("hofog(",c,") =", hofog(c))  // hofog(c)
}
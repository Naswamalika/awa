package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const NMAX int = 127 // NMAX adalah konstanta yang menyatakan batas maksimum elemen array

type tabel [NMAX]rune // tabel adalah array karakter dengan batas maksimum NMAX, rune berfungsi untuk menyimpan karakter unicode, unicode adalah standar internasional untuk encoding karakter

// Fungsi untuk mengisi array dari input satu baris
func isiArray(t *tabel, n *int, line string) {
	*n = 0 // inisialisasi jumlah elemen array, yaitu 0
	for _, char := range line { // untuk setiap karakter pada baris input
		if *n >= NMAX { // jika jumlah elemen array sudah mencapai batas maksimum
			break // hentikan loop
		}
		t[*n] = char // masukkan karakter ke dalam array
		*n++ // tambahkan jumlah elemen array
	}
}

// Fungsi untuk mencetak isi array
func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ { // untuk setiap elemen pada array
		fmt.Print(string(t[i]), " ") // cetak elemen
	}
	fmt.Println()
}

// Fungsi untuk membalikkan isi array
func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ { // untuk setiap elemen pada setengah array
		t[i], t[n-1-i] = t[n-1-i], t[i] // tukar elemen
	}
}

// Fungsi untuk memeriksa apakah array membentuk palindrom
func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ { // untuk setiap elemen pada setengah array
		if t[i] != t[n-1-i] { // jika elemen tidak sama
			return false // array tidak membentuk palindrom
		}
	}
	return true
}

func latihan4() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("-=-=- PALINDROM TEKS DALAM ARRAY -=-=-")
	fmt.Println("Masukkan teks (ketik '.' untuk berhenti):")

	for scanner.Scan() {
		line := scanner.Text()
		if strings.ToUpper(line) == "." {
			break
		}

		// Mengisi array dengan baris input
		var tab tabel
		var m int
		isiArray(&tab, &m, line)

		// Menampilkan array asli
		fmt.Print("Teks: ")
		cetakArray(tab, m)

		// Membalikkan isi array dan menampilkannya sebagai teks yang dibalik
		balikanArray(&tab, m)
		fmt.Print("Reverse Teks: ")
		cetakArray(tab, m)
		
		// Mengecek dan menampilkan hasil palindrom
		isPalindrom := palindrom(tab, m)
		fmt.Println("Palindrom:", isPalindrom)
		fmt.Println()
	}
}
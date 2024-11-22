package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv" // Package strconv berisi fungsi-fungsi untuk mengkonversi tipe data
	"strings" // Package strings berisi fungsi-fungsi untuk memanipulasi string
)

// hitungSkor menghitung jumlah soal yang diselesaikan (<= 300 menit) dan total waktu untuk soal yang diselesaikan.
func hitungSkor(soal []int, jumlahSoal *int, totalWaktu *int) { //[]int berarti array integer, dimana jumlah elemennya tidak ditentukan & *int berarti pointer ke integer
	*jumlahSoal = 0
	*totalWaktu = 0
	for _, waktu := range soal {
		// Hanya menghitung waktu soal yang diselesaikan dalam <= 300 menit
		if waktu <= 300 {
			*jumlahSoal++        // Tambah jumlah soal yang berhasil diselesaikan
			*totalWaktu += waktu // Tambah waktu ke total waktu
		}
	}
}

// soalWaktu memproses input, menghitung skor tiap peserta, dan menentukan pemenang.
func soalWaktu() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Latihan 2 Prosedur - Soal dan Waktu\n")
	// Deskripsi untuk input dari pengguna
	fmt.Println("Masukkan data peserta (format: Nama Soal1 Soal2 ... Soal8), ketik 'Selesai' untuk mengakhiri input:")

	pemenang, maxSoal, minWaktu := "", -1, 99999 // Inisialisasi variabel pemenang, -1 berasal dari jumlah soal yang diselesaikan, 99999 berasal dari total waktu

	// Memproses input hingga "Selesai" dimasukkan
	for scanner.Scan() {
		line := scanner.Text()
		if strings.ToLower(line) == "selesai" { // Memeriksa apakah pengguna mengetik 'Selesai' atau 'selesai'
			break // Keluar dari loop jika pengguna mengetik 'Selesai'
		}
		data := strings.Fields(line)  // Memisahkan baris input menjadi bagian-bagian (nama dan waktu)
		nama := data[0]               // Nama peserta
		soal := make([]int, 8)        // Array untuk menyimpan waktu pengerjaan 8 soal
		for i := 0; i < 8; i++ {      // Loop untuk mengonversi waktu soal ke integer
			soal[i], _ = strconv.Atoi(data[i+1]) // Konversi waktu soal ke integer, dan Atoi adalah fungsi untuk mengonversi string ke integer
		}
		var jumlahSoal, totalWaktu int
		hitungSkor(soal, &jumlahSoal, &totalWaktu) // Menghitung skor peserta

		// Menentukan pemenang berdasarkan jumlah soal yang diselesaikan dan total waktu
		if jumlahSoal > maxSoal || (jumlahSoal == maxSoal && totalWaktu < minWaktu) {
			pemenang, maxSoal, minWaktu = nama, jumlahSoal, totalWaktu // Jika jumlah soal yang diselesaikan lebih besar dari maxSoal atau jumlah soal yang 
			// diselesaikan sama dengan maxSoal dan total waktu lebih kecil dari minWaktu, maka pemenang, maxSoal, dan minWaktu diupdate
		}
	}

	// Output pemenang dan skornya
	fmt.Println("\n-=- Hasil Kompetisi -=-")
	fmt.Printf("Pemenang: %s\nJumlah soal yang diselesaikan: %d soal!\nTotal waktu: %d menit!\n", pemenang, maxSoal, minWaktu)
}
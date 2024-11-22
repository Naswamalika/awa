package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func latihan3() {
	var clubA, clubB string
	var winningClubs []string

	fmt.Println("\n-=-=- PERTANDINGAN BOLA -=-=-")
	fmt.Println("Masukkan nama klub A dan klub B untuk memulai pertandingan!")

	// Membaca nama klub A dan klub B
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Klub A: ")
	clubA, _ = reader.ReadString('\n') // Membaca input dari user, _ adalah variabel untuk menampung error
	clubA = strings.TrimSpace(clubA) // Menghapus spasi di awal dan akhir string, TrimSpace berfungsi untuk menghapus spasi di awal dan akhir string

	fmt.Print("Klub B: ")
	clubB, _ = reader.ReadString('\n') // Membaca input dari user, _ adalah variabel untuk menampung error
	clubB = strings.TrimSpace(clubB) // Menghapus spasi di awal dan akhir string, TrimSpace berfungsi untuk menghapus spasi di awal dan akhir string

	match := 1

	fmt.Println("\nMasukkan skor pertandingan dengan format 'skorA skorB'!")
	// Loop untuk memasukkan skor pertandingan
	for {
		fmt.Printf("Pertandingan %d: ", match)

		// Membaca skor pertandingan
		input, _ := reader.ReadString('\n') // Membaca input dari user, _ adalah variabel untuk menampung error
		input = strings.TrimSpace(input)
		scores := strings.Split(input, " ") // Memisahkan input berdasarkan spasi

		// Menghentikan loop jika input skor tidak valid
		if len(scores) != 2 { // Jika input tidak memiliki dua skor, maka hentikan loop
			fmt.Println("Input tidak valid. Pertandingan selesai!")
			break
		}

		// Mengubah string ke integer, strconv.Atoi() mengembalikan dua nilai, yaitu nilai integer dan error
		scoreA, errA := strconv.Atoi(scores[0]) 
		scoreB, errB := strconv.Atoi(scores[1]) 

		// Menghentikan loop jika skor tidak valid (negatif atau bukan angka)
		if errA != nil || errB != nil || scoreA < 0 || scoreB < 0 {
			fmt.Println("Skor tidak valid. Pertandingan selesai!")
			break
		}

		// Menentukan pemenang pertandingan dan menambahkannya ke daftar pemenang
		if scoreA > scoreB {
			winningClubs = append(winningClubs, fmt.Sprintf("Hasil %d : %s", match, clubA))
		} else if scoreA < scoreB {
			winningClubs = append(winningClubs, fmt.Sprintf("Hasil %d : %s", match, clubB))
		} else {
			winningClubs = append(winningClubs, fmt.Sprintf("Hasil %d : Draw", match))
		}

		match++
	}

	// Menampilkan review pemenang setiap pertandingan
	fmt.Println("\n-=-=- HASIL PERTANDINGAN -=-=-")
	for _, result := range winningClubs { // Untuk setiap hasil pertandingan,
		fmt.Println(result) // cetak hasil pertandingan
	}
}
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Mendefinisikan tipe data mahasiswa
type mahasiswa struct {
	nama, nim, kelas, jurusan string
	ipk                       float64
}

// Mendefinisikan array mahasiswa dengan kapasitas 2023
type arrMhs [2023]mahasiswa

// Fungsi untuk mencari indeks mahasiswa dengan IPK tertinggi
func IPK_2(T arrMhs, n int) int {
	// idx menyimpan indeks mahasiswa dengan IPK tertinggi sementara
	var idx int = 0
	for j := 1; j < n; j++ {
		if T[idx].ipk < T[j].ipk {
			idx = j
		}
	}
	return idx
}

func ne_arrayterstruktur() {
	var n int
	var data arrMhs
	scanner := bufio.NewScanner(os.Stdin)

	// Input jumlah mahasiswa
	fmt.Print("Masukkan jumlah mahasiswa (N <= 2023): ")
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())

	// Validasi jumlah mahasiswa
	if n <= 0 || n > 2023 {
		fmt.Println("Jumlah mahasiswa harus antara 1 dan 2023")
		return
	}

	// Input data mahasiswa
	for i := 0; i < n; i++ {
		fmt.Printf("\nMasukkan data mahasiswa ke-%d\n", i+1)

		// Input Nama
		fmt.Print("Nama: ")
		scanner.Scan()
		data[i].nama = scanner.Text()

		// Input NIM
		fmt.Print("NIM: ")
		scanner.Scan()
		data[i].nim = scanner.Text()

		// Input Kelas
		fmt.Print("Kelas: ")
		scanner.Scan()
		data[i].kelas = scanner.Text()

		// Input Jurusan
		fmt.Print("Jurusan: ")
		scanner.Scan()
		data[i].jurusan = scanner.Text()

		// Input IPK
		fmt.Print("IPK: ")
		scanner.Scan()
		data[i].ipk, _ = strconv.ParseFloat(scanner.Text(), 64)
	}

	// Panggil fungsi untuk mencari indeks mahasiswa dengan IPK tertinggi
	idxTertinggi := IPK_2(data, n)

	// Tampilkan data mahasiswa dengan IPK tertinggi
	fmt.Println("\nMahasiswa dengan IPK tertinggi:")
	fmt.Printf("Nama    : %s\n", data[idxTertinggi].nama)
	fmt.Printf("NIM     : %s\n", data[idxTertinggi].nim)
	fmt.Printf("Kelas   : %s\n", data[idxTertinggi].kelas)
	fmt.Printf("Jurusan : %s\n", data[idxTertinggi].jurusan)
	fmt.Printf("IPK     : %.2f\n", data[idxTertinggi].ipk)
}
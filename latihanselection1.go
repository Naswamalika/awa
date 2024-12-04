package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// fungsi selection sort pada array
func rumahKerabatSelection(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		// ini berfungsi untuk menukar posisi elemen agar terurut
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func latihanSelection1() {
	fmt.Println("-=-=-=-=- JORDAN'S HERCULES RUMAH KERABAT -=-=-=-=-")

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan jumlah daerah kerabat Hercules: ")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text()) // baris ini berfungsi untuk mengambil inputan jumlah daerah, _ berfungsi untuk mengabaikan error

	// validasi input n, n sebagai jumlah daerah kerabat Hercules
	if n <= 0 || n >= 1000 {
		fmt.Println("Huh? daerah kerabatnya Hercules harus dibawah 1000!")
		return
	}

	var results [][]int
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan jumlah (angka input pertama) & nomor-nomor rumah kerabat di daerah %d: ", i+1)
		scanner.Scan()
		line := strings.Split(scanner.Text(), " ")

		// membaca jumlah rumah kerabat di daerah, _ berfungsi untuk mengabaikan error
		m, _ := strconv.Atoi(line[0])

		// validasi input m, m sebagai jumlah rumah kerabat di daerah
		if m <= 0 || m >= 1000000 {
			fmt.Println("Banyak banget? rumah kerabatnya Hercules itu dibawah 1000000!")
			return
		}

		// konversi nomor rumah dari string ke integer
		rumah := make([]int, m)
		for j := 0; j < m; j++ {
			rumah[j], _ = strconv.Atoi(line[j+1])
		}

		rumahKerabatSelection(rumah)

		// save hasil pengurutan
		results = append(results, rumah)
	}

	// menampilkan hasil pengurutan
	fmt.Println("\n-=-=-=-=- HASIL PENGURUTAN RUMAH KERABAT MASING-MASING DAERAH -=-=-=-=-")
	for _, sortedRumah := range results {
		for i, nomor := range sortedRumah {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(nomor)
		}
		fmt.Println()
	}
}
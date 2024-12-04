package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// fungsi untuk melakukan Selection Sort
func oddAndEvenSelection(arr []int, ascending bool) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if ascending {
				if arr[j] < arr[idx] {
					idx = j
				}
			} else {
				if arr[j] > arr[idx] {
					idx = j
				}
			}
		}
		// tukar posisi elemen agar terurut
		arr[i], arr[idx] = arr[idx], arr[i]
	}
}

func latihanSelection2() {
	fmt.Println("-=-=-=-=- JORDAN'S RUMAH KERABAT HERCULES PENGURUTAN GANJIL GENAP -=-=-=-=-")

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan jumlah daerah kerabat Hercules: ")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	if n <= 0 || n >= 1000 {
		fmt.Println("Huh? Daerah kerabatnya Hercules harus di bawah 1000!")
		return
	}

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan jumlah (angka input pertama) & nomor-nomor rumah kerabat di daerah %d: ", i+1)
		scanner.Scan()
		line := strings.Split(scanner.Text(), " ")

		// membaca jumlah rumah kerabat di daerah, _ berfungsi untuk mengabaikan error
		m, _ := strconv.Atoi(line[0])

		if m <= 0 || m >= 1000000 {
			fmt.Println("Banyak banget? Rumah kerabatnya Hercules itu di bawah 1000000!")
			return
		}

		// memastikan jumlah rumah sesuai input
		if len(line)-1 != m {
			fmt.Println("Huh? Pay attention to your first number you input!")
			return
		}

		// memisahkan angka ganjil dan genap
		var oddNumbers, evenNumbers []int

		for j := 1; j <= m; j++ {
			num, _ := strconv.Atoi(line[j])
			if num%2 == 0 {
				evenNumbers = append(evenNumbers, num)
			} else {
				oddNumbers = append(oddNumbers, num)
			}
		}

		// mengurutkan angka ganjil membesar
		oddAndEvenSelection(oddNumbers, true)

		// mengurutkan angka genap mengecil
		oddAndEvenSelection(evenNumbers, false)

		// menampilkan hasil pengurutan
		fmt.Println("\n-=-=-=-=- HASIL PENGURUTAN -=-=-=-=-")
		for _, num := range oddNumbers {
			fmt.Print(num, " ")
		}
		fmt.Println(" ")
		for _, num := range evenNumbers {
			fmt.Print(num, " ")
		}
		fmt.Println(" ")
	}
}
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Fungsi untuk mengurutkan slice menggunakan Insertion Sort, SESUAI DENGAN PETUNJUK PADA SOAL NOMER 3 "SOAL LATIHAN SELECTION SORT"
func medianInsertion(data []int) {
	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1

		for j >= 0 && data[j] > key {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

// fungsi untuk menghitung median
func calculateMedian(data []int) int {
	length := len(data)
	if length%2 == 1 { // kalo panjang data hasil bagi 2 nya 1, maka panjang data ganjil dan median adalah elemen tengah
		return data[length/2]
	}
	// kalo panjang data genap, median adalah rata-rata dari dua elemen tengah
	return (data[length/2-1] + data[length/2]) / 2
}

func latihanSelection3() {
	fmt.Println("-=-=-=-=- JORDAN'S MEDIAN CALCULATOR -=-=-=-=-")
	fmt.Println("Hello hello! Pertama-tama, tolong masukkan data yang ingin dihitung mediannya dan ketik '-5313' untuk mengakhiri input!")
	fmt.Println("Oh ya, jangan lupa untuk mengetik '0' untuk menghitung median dari data saat ini!")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan datamu: ")
	scanner.Scan()
	input := scanner.Text()

	// memecah input ke dalam slice angka
	dataStrings := strings.Fields(input)
	var numbers []int

	for _, str := range dataStrings {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Huh, there's an error..", err)
			return
		}
		if num == -5313 {
			break
		}
		numbers = append(numbers, num)
	}

	var currentData []int // fungsi ini menyimpan data yang diinputkan kecuali angka 0
	zeroFound := false

	for _, num := range numbers {
		if num == 0 {
			zeroFound = true
			// fungsi ini menghitung median dari currentData, tepatnya dari angka 0 terakhir yang diinputkan
			medianInsertion(currentData)
			median := calculateMedian(currentData)
			fmt.Println("\n-=-=-=-=- HASIL MEDIAN -=-=-=-=-")
			fmt.Printf("Median dari data saat ini %v adalah: %d\n", currentData, median)
		} else {
			// menambahkan angka ke currentData kalo bukan 0
			currentData = append(currentData, num)
		}
	}

	// kalo tidak ada angka 0 dalam input, tampilkan pesan error
	if !zeroFound {
		fmt.Println("Lupa ngetik 0 ya?")
	}
}
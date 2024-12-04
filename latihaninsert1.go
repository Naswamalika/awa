package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// fungsi ini akan mengurutkan data menggunakan algoritma Insertion Sort
func jarakInsertion(data []int) {
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

// fungsi untuk memeriksa apakah data terurut memiliki jarak tetap
func checkEqualSpacing(data []int) (bool, int) {
	if len(data) < 2 {
		// kalo elemen kurang dari 2, tidak ada jarak yang dapat diperiksa
		return true, 0
	}

	distance := data[1] - data[0]
	for i := 2; i < len(data); i++ { // untuk setiap elemen data ke-i yang dimulai dari 2
		if data[i]-data[i-1] != distance { // kalo jarak antara elemen ke-i dan ke-(i-1) tidak sama dengan jarak sebelumnya
			return false, 0 // maka data tidak memiliki jarak tetap
		}
	}
	return true, distance
}

func latihanInsertion1() {
	fmt.Println("-=-=-=- JORDAN'S JARAK INSERTION SORT -=-=-=-")
	fmt.Println("Hey hey! Inputkan data berbilangan bulat positif dan akhiri dengan bilangan negatif!")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan data: ")
	scanner.Scan()
	input := scanner.Text()

	// memecah input ke dalam slice angka
	dataStrings := strings.Fields(input)
	var numbers []int

	// untuk setiap string dalam dataStrings, konversi ke int dan tambahkan ke slice numbers
	for _, str := range dataStrings {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Hmm, there's a problem.. here: ", err) // tampilkan pesan error kalo input tidak valid
			return
		}
		if num < 0 { // kalo input negatif, berhenti
			break
		}
		numbers = append(numbers, num)
	}

	// urutkan data menggunakan Insertion Sort
	jarakInsertion(numbers)

	// periksa apakah data terurut memiliki jarak tetap
	isEqual, distance := checkEqualSpacing(numbers)

	fmt.Println("\n-=-=-=-=- HASIL URUTAN JARAK -=-=-=-=-")
	fmt.Println("Data setelah diurutkan:", numbers)
	if isEqual {
		fmt.Printf("Data yang anda input berjarak %d\n", distance)
	} else {
		fmt.Println("Uhh.. ini ga tetap jaraknya.")
	}
}
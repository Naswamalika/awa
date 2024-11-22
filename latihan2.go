package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func perulanganLatihan2() {
	scanner := bufio.NewScanner(os.Stdin) // Membuat scanner baru dimana Stdin adalah input dari terminal
	var pita string
	var bungaCount int

	for {
		fmt.Printf("Bunga %d: ", bungaCount+1) // %d disini adalah format specifier untuk integer, dan bungaCount+1 adalah nilai yang akan diisi
		scanner.Scan()
		input := scanner.Text() // Mengambil input dari terminal

		if strings.ToLower(input) == "selesai" { // strings.ToLower(input) berfungsi untuk mengubah input menjadi huruf kecil semua
			break
	}

	// Jika pita masih kosong, maka pita = input, jika tidak maka pita = pita + " - " + input
	if pita == "" {
		pita = input
	} else {
		pita = pita + " - " + input
	}
	bungaCount++
	}

	fmt.Printf("Pita: %s\n", pita)
	fmt.Printf("Bunga: %d\n", bungaCount)
}
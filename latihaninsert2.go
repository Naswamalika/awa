package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

const nMax = 7919
type DaftarBuku [nMax]Buku

// Prosedur DaftarkanBuku
func DaftarkanBuku(pustaka *DaftarBuku, n *int, scanner *bufio.Scanner) {
	for i := 0; i < *n; i++ {
		fmt.Printf("\n")
		fmt.Printf("-=-=-=- BUKU KE-%d -=-=-=-\n", i+1)

		// input id
		fmt.Print("ID: ")
		scanner.Scan()
		id := strings.TrimSpace(scanner.Text())

		// input judul
		fmt.Print("Judul: ")
		scanner.Scan()
		judul := strings.TrimSpace(scanner.Text())

		// input penulis
		fmt.Print("Penulis: ")
		scanner.Scan()
		penulis := strings.TrimSpace(scanner.Text())

		// input penerbit
		fmt.Print("Penerbit: ")
		scanner.Scan()
		penerbit := strings.TrimSpace(scanner.Text())

		var eksemplar, tahun, rating int
		var err error

		// untuk eksemplar, validasi input harus angka
		for {
			fmt.Print("Eksemplar: ")
			scanner.Scan()
			eksemplar, err = strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err == nil {
				break
			}
			fmt.Println("Kamu barusan input non-angka ya? Do it AGAIN.")
		}

		// untuk tahun, validasi input harus angka
		for {
			fmt.Print("Tahun: ")
			scanner.Scan()
			tahun, err = strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err == nil {
				break
			}
			fmt.Println("Kamu barusan input non-angka ya? Do it AGAIN.")
		}

		// untuk setiap rating, validasi input harus angka
		for {
			fmt.Print("Rating: ")
			scanner.Scan()
			rating, err = strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err == nil {
				break
			}
			fmt.Println("Kamu barusan input non-angka ya? Do it AGAIN.")
		}

		// memasukkan data buku ke dalam array pustaka
		pustaka[i] = Buku{
			id:       id,
			judul:    judul,
			penulis:  penulis,
			penerbit: penerbit,
			eksemplar: eksemplar,
			tahun:     tahun,
			rating:    rating,
		}
	}
}

// prosedur CetakTerfavorit untuk mencetak buku terfavorit
func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n == 0 {
		fmt.Println("?? Pustakanya kosong??")
		return
	}
	terfavorit := pustaka[0]
	for i := 1; i < n; i++ {
		if pustaka[i].rating > terfavorit.rating {
			terfavorit = pustaka[i]
		}
	}
	fmt.Println("\n-=-=-=- BUKU TERFAVORIT -=-=-=-")
	fmt.Printf("Judul: %s\n", terfavorit.judul)
	fmt.Printf("Penulis: %s\n", terfavorit.penulis)
	fmt.Printf("Penerbit: %s\n", terfavorit.penerbit)
	fmt.Printf("Tahun: %d\n", terfavorit.tahun)
}

// prosedur UrutBuku untuk mengurutkan buku berdasarkan rating menurun dengan insertion sort
func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		key := pustaka[i]
		j := i - 1
		for j >= 0 && pustaka[j].rating < key.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key 
	}
}

// prosedur Cetak5Terbaru untuk mencetak 5 buku dengan rating tertinggi
func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	fmt.Println("\n-=-=-=- 5 BUKU DENGAN RATING TERTINGGI -=-=-=-")
	for i := 0; i < 5 && i < n; i++ {
		fmt.Printf("%d. %s\n", i+1, pustaka[i].judul)
	}
}

// prosedur CariBuku untuk mencari buku dengan rating tertentu
func CariBuku(pustaka DaftarBuku, n int, r int) {
	low, high := 0, n-1
	for low <= high {
		mid := (low + high) / 2
		if pustaka[mid].rating == r {
			buku := pustaka[mid]
			fmt.Printf("\n-=-=-=- BUKU DENGAN RATING %d -=-=-=-\n", r)
			fmt.Printf("Judul: %s\n", buku.judul)
			fmt.Printf("Penulis: %s\n", buku.penulis)
			fmt.Printf("Penerbit: %s\n", buku.penerbit)
			fmt.Printf("Tahun: %d\n", buku.tahun)
			fmt.Printf("Eksemplar: %d\n", buku.eksemplar)
			fmt.Printf("Rating: %d\n", buku.rating)
			return
		} else if pustaka[mid].rating < r {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	fmt.Println("Tidak ada buku dengan rating seperti itu.")
}

func latihanInsertion2() {
	fmt.Println("-=-=-=- JORDAN'S PUSTAKA BUKU PROGRAM -=-=-=-")

	var pustaka DaftarBuku
	var nPustaka int

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Hi! Go ahead masukkan jumlah buku di pustaka: ")
	scanner.Scan()
	input := scanner.Text()
	nPustaka, err := strconv.Atoi(input)

	// validasi input jumlah buku antara 1 dan nMax, yaitu 7919
	if err != nil || nPustaka <= 0 || nPustaka > nMax {
		fmt.Printf("Uhh jumlah buku harus antara 1 dan %d.\n", nMax)
		return
	}

	// memanggil prosedur DaftarkanBuku untuk mengisi data buku
	DaftarkanBuku(&pustaka, &nPustaka, scanner)

	// memanggil prosedur CetakTerfavorit untuk mencetak buku terfavorit
	CetakTerfavorit(pustaka, nPustaka)

	// memanggil prosedur UrutBuku untuk mengurutkan buku berdasarkan rating
	UrutBuku(&pustaka, nPustaka)

	// memanggil prosedur Cetak5Terbaru untuk mencetak 5 buku dengan rating tertinggi
	Cetak5Terbaru(pustaka, nPustaka)

	// mencari buku dengan rating yang diinput
	var r int
	fmt.Println("\n-=-=-=- CARI BUKU DARI RATING -=-=-=-")
	fmt.Print("Dari sekian banyak buku, inputlah rating buku yang mau anda cari: ")
	scanner.Scan()
	r, err = strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("?? It ain't there.")
		return
	}
	CariBuku(pustaka, nPustaka, r) // memanggil prosedur CariBuku untuk mencari buku dengan rating yang diinput
}
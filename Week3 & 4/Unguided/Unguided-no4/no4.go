package main

import "fmt"

const phi = 3.14

// Procedure hitung persegi
func hitungPersegi(sisi int) {
	luas := sisi * sisi
	keliling := 4 * sisi

	fmt.Println()
	fmt.Println("Masukkan sisi : ", sisi)
	fmt.Println("Luas persegi : ", luas)
	fmt.Println("Keliling persegi : ", keliling)
}

// Procedure hitung persegi panjang
func hitungPersegiPanjang(panjang, lebar int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)

	fmt.Println()
	fmt.Println("Masukkan panjang : ", panjang)
	fmt.Println("Masukkan lebar : ", lebar)
	fmt.Println("Luas persegi panjang : ", luas)
	fmt.Println("Keliling persegi panjang : ", keliling)
}

// Procedure hitung lingkaran
func hitungLingkaran(jariJari float64) {
	luas := phi * jariJari * jariJari
	keliling := 2 * phi * jariJari

	fmt.Println()
	fmt.Println("Masukkan jari-jari : ", jariJari)
	fmt.Println("Luas lingkaran : ", luas)
	fmt.Println("Keliling lingkaran : ", keliling)
}

func main() {
	var pilihan int

	fmt.Println("--- PROGRAM BANGUN DATAR ---")
	fmt.Println("1. Hitung luas & keliling persegi")
	fmt.Println("2. Hitung luas & keliling persegi panjang")
	fmt.Println("3. Hitung luas & keliling lingkaran")
	fmt.Print("Pilihan : ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		hitungPersegi(1057)
	case 2:
		hitungPersegiPanjang(106, 88)
	case 3:
		hitungLingkaran(13.87)
	default:
		fmt.Println("Pilihan tidak valid")
	}
}
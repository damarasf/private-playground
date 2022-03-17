package main

import (
	"fmt"
)

// Check Point:
// Two Dimention Area
// - Input: Select Choice
// - Output: Result Operation

// Contoh:
// Input:
// 1: Rectange Area
// 2: Rectangular Area
// 3: Triangle Area
// 4: Circle Area
// - Enter choice 1, 2, 3, or 4: 1 | 2 | 3 | 4
//   - (1) Rectange Area
//   	- Masukkan sisi : 5
//   - (2) Rectangular Area
// 		- Masukkan panjang : 5
// 		- Masukkan lebar : 10
// 	 - (3) Triangle Area
// 		- Masukkan panjang alas segitiga: 5
// 		- Masukkan tinggi segitiga: 10
// 	 - (4) Circle Area
//      - Masukkan jari-jari : 4

// Output:
// - (1) Luas Persegi adalah : 25
// - (2) Luas Persegi Panjang adalah : 50
// - (3) Luas Segitiga adalah : 25
// - (4) Luas Lingkaran adalah : 50.240000

func main() {
	// TODO: answer here
	var sisi, panjang, lebar, palas, tsegitiga, r float32
	var pilih int = 0
	var hasil float32 = 0

	fmt.Println("1: Rectangle Area")
	fmt.Println("2: Rectangular Area")
	fmt.Println("3: Triangle Area")
	fmt.Println("4: Circle Area")
	fmt.Print("Enter choice: ")
	fmt.Scan(&pilih)

	switch pilih {
	case 1:
		fmt.Println ("(1) Rectange Area")
		fmt.Printf ("Masukkan Sisi : ")
		fmt.Scan(&sisi)
		hasil = sisi * sisi
		fmt.Printf("Luas Persegi adalah: %f", hasil)
	case 2:
		fmt.Println ("(2) Rectangular Area")
		fmt.Printf ("Masukkan Panjang : ")
		fmt.Scan(&panjang)
		fmt.Printf ("Masukkan Lebar : ")
		fmt.Scan(&lebar)
		hasil = panjang * lebar
		fmt.Printf("Luas Persegi Panjang: %f", hasil)
	case 3:
		fmt.Println ("(3) Triangle Area")
		fmt.Printf ("Masukkan Alas : ")
		fmt.Scan(&palas)
		fmt.Printf ("Masukkan Tinggi : ")
		fmt.Scan(&tsegitiga)
		hasil = palas * tsegitiga / 2
		fmt.Printf("Luas Segitiga adalah: %f", hasil)
	case 4:
		const pi = 3.14
		fmt.Println ("(4) Circle Area")
		fmt.Printf ("Masukkan Jari - Jari : ")
		fmt.Scan(&r)
		hasil = pi * r * r
		fmt.Printf("Luas Lingkaran adalah: %f", hasil)
	default:
		fmt.Println("Invalid choice")
	}
}

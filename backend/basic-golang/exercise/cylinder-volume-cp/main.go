package main

import (
	"fmt"
)

// Check Point:
// Menghitung volume tabung
// - Input: jari-jari, tinggi
// - Output: volume tabung

// Contoh:
// Input:
// - Masukkan jari-jari alas tabung: 2
// - Masukkan tinggi tabung : 20
// Output:
// - Jadi volumenya adalah : 251.200012

func main() {
	// TODO: answer here
	var (
		jari, tinggi, area float32
	)
	const pi = 3.14
	fmt.Printf("Masukkan jari - jari : ")
	fmt.Scan(&jari)
	fmt.Printf("Masukkan tinggi : ")
	fmt.Scan(&tinggi)
	area = pi * jari * jari * tinggi
	fmt.Printf("Jadi volumenya adalah : %f", area)
}

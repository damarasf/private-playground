package main

import (
	"fmt"
	"sort"
)

type Mahasiswa struct {
	Nama     string
	NPM      int
	Angkatan int
}

type Mahasiswas []Mahasiswa

func (s Mahasiswas) Len() int {
	return len(s)
}

func (s Mahasiswas) Less(i, j int) bool {
	swapped := false
	if s[i].NPM > s[j].NPM {
		swapped = true
	}
	if s[i].NPM == s[j].NPM {
		if s[i].Angkatan > s[j].Angkatan {
			swapped = true
		}
	}
	return swapped
}

func (s Mahasiswas) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Mahasiswas) TopStudents() []string {
	sort.Sort(s)
	names := []string{}
	for _, score := range s {
		names = append(names, score.Nama)
	}
	return names
}

func main() {
	// User bisa input
	/*
		1. Input Jumlah Data
		2. Input Nama
		3. Input NPM
		4. Input Angkatan

		jumlah data = 5
			Agus 123 2019
			Ade 321 2020
			Hendra 345 2021
			Yudi 124 2020
			Yuli 432 2019

		Output: [nama: Agus,  NPM: 123, Angkatan: 2019
			Ade 321 2020
			Hendra 345 2021
			Yudi 124 2020
			Yuli 432 2019]
	*/

	var choose int
	fmt.Println("1. Input Jumlah Data")
	fmt.Println("2. Input Nama")
	fmt.Println("3. Input NPM")
	fmt.Println("4. Input Angkatan")
	fmt.Println("5. Exit")
	fmt.Println("Pilih : ")
	fmt.Scanln(&choose)

	data := Mahasiswas([]Mahasiswa{
		{"Agus", 123, 2019},
		{"Ade", 321, 2020},
		{"Hendra", 345, 2021},
		{"Yudi", 124, 2020},
		{"Yuli", 432, 2019},
	})

	switch choose {
	case 1:
		var jumlah int
		fmt.Println("Jumlah Data : ")
		fmt.Scanln(&jumlah)
		for i := 0; i < jumlah; i++ {
			fmt.Println("Nama : ")
			var nama string
			fmt.Scanln(&nama)
			fmt.Println("NPM : ")
			var npm int
			fmt.Scanln(&npm)
			fmt.Println("Angkatan : ")
			var angkatan int
			fmt.Scanln(&angkatan)
			data = append(data, Mahasiswa{nama, npm, angkatan})
		}
		fmt.Println(data)
	case 2:
		var nama string
		fmt.Println("Nama : ")
		fmt.Scanln(&nama)
		fmt.Println("NPM : ")
		var npm int
		fmt.Scanln(&npm)
		fmt.Println("Angkatan : ")
		var angkatan int
		fmt.Scanln(&angkatan)
		data = append(data, Mahasiswa{nama, npm, angkatan})
		fmt.Println(data)
	case 3:
		var npm int
		fmt.Println("NPM : ")
		fmt.Scanln(&npm)
		fmt.Println("Angkatan : ")
		var angkatan int
		fmt.Scanln(&angkatan)
		data = append(data, Mahasiswa{"", npm, angkatan})
		fmt.Println(data)
	case 4:
		var angkatan int
		fmt.Println("Angkatan : ")
		fmt.Scanln(&angkatan)
		data = append(data, Mahasiswa{"", 0, angkatan})
		fmt.Println(data)
	case 5:
		fmt.Println("Exit")
		break
	default:
		fmt.Println("Pilih yang benar")
	}

	sort.Sort(data)
	fmt.Println(data.TopStudents())
}

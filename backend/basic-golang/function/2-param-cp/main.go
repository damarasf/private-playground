package main

import "fmt"

//memanggil fungsi goodMorning
//fungsi goodMorning akan melakukan print selamat pagi + name yang didapat dari parameter fungsi
func main() {
	goodMorning("Damara")
	goodMorning("Anto")

}

// TODO: answer here
func goodMorning(name string) {
	fmt.Println("Hello Good Morning", name)
}

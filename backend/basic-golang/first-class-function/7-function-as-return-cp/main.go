package main

import "fmt"

func main() {
	// fungsi ini akan mengembalikan fungsi berikut
	// func(x, y int) int {
	// 	return x * y
	// }
	// TODO: answer here
	getAreaFunc := func(x int, y int) func(int) int {
		return func( int) int {}
	}
	areaF := getAreaFunc()
	res := areaF(3, 4) // 12

	fmt.Println(res)
}

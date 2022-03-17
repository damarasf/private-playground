package main

import "fmt"

// Check Point:
// Diamond Pattern
// - Input: Size: 5
// - Output:
//         *
//        ***
//       *****
//      *******
//     *********
//    ***********
//     *********
//      *******
//       *****
//        ***
//         *

func main() {
	var i, j, k, size int

	fmt.Printf("Enter Size : ")
	fmt.Scan(&size)

	// TODO: answer here
	for i = 1; i <= size; i++ {
		for j = 1; j <= size-i; j++ {
			fmt.Printf(" ")
		}
		for k = 1; k <= i*2-1; k++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}

	for i = size - 1; i > 0; i-- {
		for j = 1; j <= size-i; j++ {
			fmt.Printf(" ")
		}
		for k = 1; k <= i*2-1; k++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

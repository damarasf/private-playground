package main

import "fmt"

func main() {
	// arr := []int{0, 1, 2, 3}
	arr := []int{12, 5, 10, 1, 3}

	res := SumAll(arr)
	fmt.Println(res)

	// Try correct answer:
	resCorrect := SumAllCorrect(arr)
	fmt.Println(resCorrect)
}

func SumAll(arr []int) int {
	res := 0
	for val := range arr {
		res += val
	}
	return res
}

func SumAllCorrect(arr []int) int {
	resCorrect := 0
	for _, val := range arr {
		resCorrect += val
	}

	if len(arr) == 0 {
		return 0
	}

	return resCorrect
}

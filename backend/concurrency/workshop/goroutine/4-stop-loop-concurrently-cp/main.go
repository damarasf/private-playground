package main

import (
	"fmt"
	"time"
)

//bagaiman cara menghentikan loop ?
func stopLoop(loop *bool) {
	// TODO: answer here
	a := 0
	for a < 10 {
		fmt.Println("loop")
		a++
	}

	*loop = false

	time.Sleep(200 * time.Millisecond)

	fmt.Println("main stop")
}

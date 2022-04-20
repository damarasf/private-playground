package main

import (
	"time"
)

//bagaiman cara menghentikan loop ?
func stopLoop(loop *bool) {
	// TODO: answer here
	*loop = false
	time.Sleep(200 * time.Millisecond)
}

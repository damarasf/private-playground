package main

import (
	"fmt"
	"time"
)

type character struct {
	name            string
	activity        string
	defaultActivity string
}

func (c *character) awake(movementInput, attackInput chan string) {
	for {
		//sama seperti checkpoint sebelumnya
		//tetapi sekarang menambahkan default activity yang menjalankan
		//fmt.Printf("%s %s\n", c.name, c.defaultActivity)
		//time.Sleep(100 * time.Millisecond)
		// TODO: answer here
		select {
		case movement := <-movementInput:
			fmt.Printf("%s bergerak ke %s\n", c.name, movement)
			c.activity = movement
		case attack := <-attackInput:
			fmt.Printf("%s melakukan serangan %s\n", c.name, attack)
			c.activity = attack
		default:
			fmt.Printf("%s %s\n", c.name, c.defaultActivity)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

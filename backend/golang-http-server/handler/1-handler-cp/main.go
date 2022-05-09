package main

import (
	"fmt"
	"net/http"
	"time"
)

// Dari contoh yang diberikan, buatlah sebuah handler dengan menggunakan HandlerFunc yang menampilkan nama hari, bulan, dan tahun.
// Hint, gunakan time.Weekday, time.Day, time.Month, dan time.Year

func GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) { // TODO: replace this
		t := time.Now()
		expected := fmt.Sprintf("%v, %v %v %v", t.Weekday(), t.Day(), t.Month(), t.Year())
		fmt.Fprint(writer, expected)
	}
}

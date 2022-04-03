package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"
)

//Buat function untuk menghitung average score siswa
//panggil function didalam template

type Student struct {
	Name   string
	Scores []float64
}

func (s Student) CalculateScore(scores []float64) float64 {
	// TODO: answer here
	var sum float64
	for _, score := range scores {
		sum += score
	}
	return sum / float64(len(scores))
}

func (s Student) GenerateStudentTemplate() string {
	buff := new(bytes.Buffer)
	// TODO: answer here
	tmp1 := template.New("Template_1")
	tmp1, err := tmp1.Parse(`Hello {{.Name}}, Nilai rata-rata kamu {{.CalculateScore .Scores}}`)
	if err != nil {
		log.Fatalf("parse error: %s", err.Error())
	}
	err = tmp1.Execute(buff, s)
	if err != nil {
		log.Fatalf("execute template error: %s", err.Error())
	}
	return buff.String()
}

func NewStudent(name string, scores []float64) Student {
	return Student{name, scores}
}

// main function
func main() {
	std := NewStudent("Rogu", []float64{10, 11, 12})
	fmt.Println(std.GenerateStudentTemplate())
}

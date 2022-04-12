package main

import (
	"fmt"
	"sort"
)

// # Scoreboard

// Dalam penilaian ujian, nilai akhir adalah hasil dari perhitungan berikut:

// Nilai = 4 x Jumlah Benar - 1 x Jumlah Salah`

// Kita ingin mengurutkan hasil ujian siswa, urut berdasarkan nilai tertinggi ke terendah.
// Jika ada yang nilainya sama, maka:
// Yang `Jumlah Benar`-nya lebih tinggi akan diurutkan di atas.
// Jika masih sama:
// Yang `Nama`-nya lebih awal akan diurutkan di atas

type Score struct {
	Name    string
	Correct int
	Wrong   int
	Empty   int
}
type Scores []Score

func (s Scores) Len() int {
	return len(s)
}

func (s Scores) Less(i, j int) bool {
	swapped := false
	if s[i].Correct*4-s[i].Wrong*1 > s[j].Correct*4-s[j].Wrong*1 {
		swapped = true
	}
	if s[i].Correct*4-s[i].Wrong*1 == s[j].Correct*4-s[j].Wrong*1 {
		if s[i].Correct > s[j].Correct {
			swapped = true
		}
	}
	return swapped
}

func (s Scores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Scores) TopStudents() []string {
	sort.Sort(s)
	names := []string{}
	for _, score := range s {
		names = append(names, score.Name)
	}
	return names
}

func main() {
	scores := Scores([]Score{
		{"Levi", 3, 2, 2},
		{"Agus", 3, 4, 0},
		{"Doni", 3, 0, 7},
		{"Ega", 3, 0, 7},
		{"Anton", 2, 0, 5},
	})
	sort.Sort(scores)
	fmt.Println(scores.TopStudents())
}

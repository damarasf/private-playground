// Tulis program sebagai fungsi yang memeriksa apakah dua set dikatakan intersection atau tidak.
// Jika kedua set intersection, fungsi tersebut harus mengembalikan nilai intersection nya.
//
// Contoh 1
// Input: a = {"Java", "Python", "Javascript", "C ++", "C#"}, b = {"Swift", "Java", "Kotlin", "Python"}
// Output: {"Java", "Python"}
// Explanation: intersection dari a dan b adalah "Java" dan "Python"
//
// Contoh 2
// Input: a = {"Java", "Python"}, b = {"Swift"}
// Output: {}
// Explanation: tidak ada intersection dari a dan b

package main

import "fmt"

func main() {
	var str1 = []string{"Java", "Python", "Javascript", "C ++", "C#"}
	var str2 = []string{"Swift", "Java", "Kotlin", "Python"}
	fmt.Println(Intersection(str1, str2))
}

func Intersection(str1, str2 []string) (inter []string) {
	// hash := make(map[string]bool)

	// for _, item := range str1 {
	// 	hash[item] = true
	// }

	// for _, item := range str2 {
	// 	if _, exist := hash[item]; exist {
	// 		inter = append(inter, item)
	// 	}
	// }
	// return
	var str1_nodups = RemoveDuplicates(str1)
	var str2_nodups = RemoveDuplicates(str2)

	for _, item := range str1_nodups {
		for _, item2 := range str2_nodups {
			if item == item2 {
				inter = append(inter, item)
			}
		}
	}
	return
}

func RemoveDuplicates(elements []string) (nodups []string) {
	hash := map[string]bool{}

	for v := range elements {
		if hash[elements[v]] == true {
			// Do not add duplicate.
		} else {
			hash[elements[v]] = true
			nodups = append(nodups, elements[v])
		}
	}
	return
}

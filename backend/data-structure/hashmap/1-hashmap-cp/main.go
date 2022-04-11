package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	people := []Person{{name: "Bob", age: 21}, {name: "Sam", age: 28}, {name: "Ann", age: 21}, {name: "Joe", age: 22}, {name: "Ben", age: 28}}
	fmt.Println(AgeDistribution(people))
	fmt.Println(FilterByAge(people, 21))
}

func AgeDistribution(people []Person) map[int]int {
	if len(people) == 0 {
		return nil
	}

	distribution := make(map[int]int)
	for _, person := range people {
		distribution[person.age]++
	}

	return distribution
}

func FilterByAge(people []Person, age int) []Person {
	if len(people) == 0 {
		return nil
	}

	filteredPeople := make([]Person, 0)
	for _, person := range people {
		if person.age == age {
			filteredPeople = append(filteredPeople, person)
		}
	}

	return filteredPeople
}

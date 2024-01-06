package main

import (
	"fmt"
)

func hitungLuasSegitiga(alas float32, tinggi float32) float32 {
	return (alas * tinggi) / 2
}

func getFullName() (string, string, int) {
	return "Adit", "Nugroho", 90
}

func sumAll(numbers ...int) int {
	total := 0

	for _, num := range numbers {
		total += num
	}

	return total
}

func sayHelloWithFilter(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func badWordFilter(name string) string {
	if name == "anjing" || name == "babi" {
		return "*****"
	} else {
		return name
	}
}

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	luasSegitiga := hitungLuasSegitiga(10, 5)
	fmt.Println(luasSegitiga)

	firstName, lastName, num := getFullName()
	fmt.Println(firstName, lastName, num)

	fmt.Println(sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(sumAll(nums...))

	hitungSemua := sumAll

	fmt.Println(hitungSemua(1, 3, 5, 7, 9))

	sayHelloWithFilter("Adit", badWordFilter)
	sayHelloWithFilter("anjing", badWordFilter)

	registerUser("didit", func(name string) bool {
		return name == "wedus"
	})
}

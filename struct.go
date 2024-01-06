package main

import (
	"fmt"
)

type Customer struct {
	id            int
	name, address string
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My name is", customer.name)
}

func (customer Customer) destination(place string) {
	fmt.Println("let's go to", place, " from", customer.address)
}

func ageValidate(age int) bool {
	if age == 18 {
		return true
	} else {
		return false
	}
}

func endApp() {
	fmt.Println("End Application!")
	msg := recover()
	fmt.Println("terjadi error :", msg)
}

func (customer Customer) guessAge(age int, err bool) {
	defer endApp()
	fmt.Println("i guess your age is", age)
	if err {
		panic("ERROR APP")
	}
	if ageValidate(age) {
		fmt.Println("you correct!")
	} else {
		fmt.Println("you wrong!")
	}

}
func main() {

	adit := Customer{
		id:      1,
		name:    "Adit",
		address: "Jl. Raya",
	}

	adit.sayHello("yanto")
	fmt.Println(adit)

	dodi := Customer{2, "dodi kurniawan", "jalan raya"}
	dodi.destination("binjai")

	alex := Customer{4, "alex ginting", "medan"}

	alex.guessAge(20, false)
}

package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
	age  int
}

type Vehicle struct {
	Name string
	fuel string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

func WhatVehicle(value HasName) {
	fmt.Println("Your Vehicle is", value.GetName())
}

func (vhc Vehicle) GetName() string {
	return vhc.Name
}

func (person Person) GetName() string {
	return person.Name
}

func Ups() any {
	//return 1
	return "ups"
}

func main() {
	budi := Person{Name: "budi hartono", age: 19}
	SayHello(budi)

	porsche := Vehicle{Name: "Porsche 11", fuel: "Solar"}
	anyy := Ups()
	fmt.Println(anyy)
	WhatVehicle(porsche)

}

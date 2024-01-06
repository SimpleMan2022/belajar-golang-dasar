package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr." + man.Name
}

func main() {
	adit := Man{"Adit Nugroho"}
	adit.Married()
	fmt.Println(adit.Name)
}

package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeCountryToIndo(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address1 := Address{}
	changeCountryToIndo(&address1)
	fmt.Println(address1)
}

package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	//address1 := Address{"Medan", "Sumatera Utara", "Indonesia"}
	//address2 := address1
	//
	//address2.City = "Binjai"
	//address2 := &address1
	//
	//address2.City = "Binjai"
	//fmt.Println(address1)
	//fmt.Println(address2)
	//
	//*address2 = Address{"jakarta", "DKI Jakarta", "Indonesia"}
	//fmt.Println(address1)
	//fmt.Println(address2)
	address3 := new(Address)
	address4 := address3

	address4.Country = "Indonesia"
	fmt.Println(address3)
	fmt.Println(address4)

}

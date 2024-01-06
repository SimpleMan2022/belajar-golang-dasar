package main

import "fmt"

func main() {

	numDays := 2

	switch numDays {
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("Monday")
	case 3:
		fmt.Println("Tuesday")
	default:
		fmt.Println("inputan salah!")
	}

	switch length := len("adits"); length >= 5 {
	case true:
		fmt.Println("terlalu panjang")
	case false:
		fmt.Println("kurang dari 5")
	}

	switch {
	case numDays == 1:
		fmt.Println("Sunday")
	case numDays == 2:
		fmt.Println("Monday")
	}
}

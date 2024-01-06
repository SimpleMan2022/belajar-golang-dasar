package main

import "fmt"

func main() {
	value := 6

	if isEven := value%2 == 0; isEven {
		fmt.Println("Value is even:", isEven) // cetak nilai genap
	} else {
		fmt.Println("Value is odd")
	}
}

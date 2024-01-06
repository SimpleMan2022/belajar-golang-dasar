package main

import (
	"fmt"
)

func main() {

	type noKtp string

	var nomorKtp noKtp = "12345"
	nomorNIK := noKtp("12345678")
	fmt.Println(nomorKtp)
	fmt.Println(nomorNIK)
}

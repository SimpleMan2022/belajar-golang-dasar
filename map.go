package main

import (
	"fmt"
)

func main() {
	orang1 := map[string]string{
		"name": "Adit",
		"age":  "19",
		"Ipk":  "4.0",
	}

	fmt.Println(orang1["name"])
	fmt.Println(len(orang1))
	orang1["name"] = "didit"
	fmt.Println(orang1["name"])

	orang2 := make(map[string]string)

	orang2["name"] = "bunga"
	orang2["age"] = "19"

	fmt.Println(orang2)
	fmt.Println(orang2)
}

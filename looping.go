package main

import (
	"fmt"
)

func main() {

	// count := 5
	// for i := 0; i < count; i++ {
	// 	fmt.Println("index : ", i)
	// }

	// for count <= 10 {
	// 	fmt.Println("index : ", count)
	// 	count++
	// }

	names := [...]string{"Adit", "Nugroho", "Bayu", "agnes", "jessica"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// for index, name := range names {
	// 	fmt.Println(index, name)
	// }

	// for _, name := range names {
	// 	fmt.Println(name)
	// }

	// break

	// for index, name := range names {
	// 	if index == 2 {
	// 		break
	// 	}

	// 	fmt.Println(index, name)
	// }

	// continue

	for index, name := range names {
		if name == "agnes" {
			continue
		}

		fmt.Println(index, name)
	}
}

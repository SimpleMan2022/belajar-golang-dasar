package main

import "fmt"

func random() any {
	return true
}

func numS() any {
	return 1
}

func main() {
	result := random()
	//resultString := result.(string)
	//fmt.Println(resultString)
	//
	//numbers := numS()
	//parseNum := numbers.(int)
	//fmt.Println(parseNum)

	switch value := result.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	case bool:
		fmt.Println("boolean", value)
	default:
		fmt.Println("Unknown", "value")
	}
}

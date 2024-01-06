package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	orang1 := newMap("adit nugroho")
	orang2 := newMap("")
	if orang2 == nil {
		fmt.Println("datanya kosong")
	} else {
		fmt.Println(orang2)
	}
	fmt.Println(orang1)
}

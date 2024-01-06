package main

import "fmt"

func main() {
	// var nilaiUas [2]int

	// nilaiUas[0] = 90
	// nilaiUas[1] = 95

	// tinggiBadan := [3]int {170, 180, 190}

	// fmt.Println("Nilai UAS pertama : ", nilaiUas[0])
	// fmt.Println("Nilai UAS terakhir : ", nilaiUas[1])

	// fmt.Println(nilaiUas)
	// nilaiUas[0] = 100

	// fmt.Println("Nilai UAS pertama : ", nilaiUas[0])
	// fmt.Println(tinggiBadan)
	// fmt.Println(len(nilaiUas))

	nilaiBData := [...]int{90, 95, 100, 84, 75, 86}
	// namaMhs := [...] string {"bayu", "adit", "nugroho"}
	// fmt.Println(nilaiBData)
	// fmt.Println(namaMhs)

	// var slice []string = namaMhs[1:]
	// fmt.Println(slice)
	// fmt.Println(len(slice))
	// fmt.Println(cap(slice))

	nilaiBaru := nilaiBData[3:]
	nilaiBaru[0] = 100
	nilaiBaru[2] = 60
	fmt.Println(nilaiBData)

	nilaiBaru2 := append(nilaiBaru, 50)
	fmt.Println(nilaiBaru2)
	fmt.Println(nilaiBData)

	// make

}

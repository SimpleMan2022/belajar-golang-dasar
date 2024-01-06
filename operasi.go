package main

import "fmt"

func main() {
	// operasi matematika
	// var a = 10
	// // b := 20
	// // c := a + b
	// // d := a - b
	// // e := a * b
	// a += 10
	// a := a++
	// a++
	// // fmt.Println(c)
	// // fmt.Println(d)
	// // fmt.Println(e)
	// fmt.Println(a)

	//operasi perbandingan
	// name1 := "Adit"
	// name2 := "Adit"

	// //membuat nama menjad huruh kecil semua
	// var hasil = name1 != name2

	// fmt.Println(hasil)

	//operasi logika

	nilaiUts := 90
	nilaiUas := 95
	var lulusUts bool = nilaiUts > 86
	var lulusUas bool = nilaiUas > 96

	naikKelas := lulusUts && lulusUas

	fmt.Println(naikKelas)
}

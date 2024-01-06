package main

import "fmt"

func main() {

	name := "Adit nugroho"
	usia := 15
	isMakeKtp := false

	if usia >= 17 && isMakeKtp {
		fmt.Println("Anda boleh memberikan suara!")
	} else if usia >= 17 && !isMakeKtp {
		fmt.Println("silahkan buat ktp terlebih dahulu!")
	} else {
		fmt.Println("maaf, anda belum cukup umur")
	}

	if length := len(name); length < 5 {
		fmt.Println("nama terlalu pendek!")
	} else {
		fmt.Println("selamat datang", name)
	}

	value := 6

	if evenNum, isEven := value%2 == 0; isEven {
		fmt.Println("Value is even:", evenNum)
	} else {
		fmt.Println("Value is odd")
	}

}

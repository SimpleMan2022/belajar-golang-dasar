package main

import "fmt"

func logging() {
	fmt.Println("Selesai")
}

func startApplication() {
	defer logging()
	fmt.Println("Start")
}

func endApplication() {
	fmt.Println("Program selesai!")
	msg := recover()
	fmt.Println("terjadi error", msg)
}

func startApp(err bool) {
	defer endApplication()

	if err {
		panic("APLIKASI ERROR")
	}

	fmt.Println("Start App")

}

func main() {
	// startApplication()
	startApp(true)
}

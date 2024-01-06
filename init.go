package main

import (
	"belajar-go/database"
	_ "belajar-go/internal"
	"fmt"
)

func main() {
	db := database.GetDatabase()
	fmt.Println(db)
}

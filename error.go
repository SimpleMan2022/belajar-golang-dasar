package main

import (
	"errors"
	"fmt"
)

// custom error

type validationError struct {
	Message string
}

type notFoundError struct {
	Message string
}

func (validation *validationError) Error() string {
	return validation.Message

}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"validation error"}
	}

	if id != "adit" {
		return &notFoundError{"data not found"}
	}

	return nil
}

func pembagian(nilai, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan nol")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	//result, err := pembagian(20, 0)
	//if err == nil {
	//	fmt.Println("hasil :", result)
	//} else {
	//	fmt.Println("error : ", err.Error())
	//
	//}

	save := SaveData("s", nil)

	if save != nil {
		//if validationErr, ok := save.(*validationError); ok {
		//	fmt.Println("validasi error : ", validationErr)
		//} else if notFoundErr, ok := save.(*notFoundError); ok {
		//	fmt.Println("not found error:", notFoundErr)
		//} else {
		//	fmt.Println("unknown error:", save.Error())
		//}

		switch finalError := save.(type) {
		case *validationError:
			fmt.Println("validasi error : ", finalError.Error())
		case *notFoundError:
			fmt.Println("not found  error : ", finalError.Error())
		default:
			fmt.Println("unknown error : ", finalError.Error())
		}
	} else {
		//sukses
		fmt.Println("sukses")
	}

}

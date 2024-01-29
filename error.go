package main

import (
	"fmt"
)

func main() {
	value := 130
	err := fmt.Errorf("Error: nilai yang dimasukan %d terlalu besar", value)
	fmt.Println(err)
}

// func main() {
// 	err := errors.New("Ini adalah kesalahan sederhana")
// 	fmt.Println(err)
// }

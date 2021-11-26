package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is pointers examples")

	//pointer declaration
	x := 10
	//var pointerToX *int
	pointerToX := &x
	fmt.Println(pointerToX)  // prints memory address - 0xc000012088
	fmt.Println(*pointerToX) // prints value to that memory address - 10
	z := 5 + *pointerToX
	fmt.Println(z) // 15

}

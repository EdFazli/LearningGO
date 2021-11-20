package main

import (
	"fmt"
)

//using struct to simulate named parameters
type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	fmt.Println("This is functions examples")
	result := div(5, 2)
	fmt.Println(result) // 2

	//using struct to simulate named parameters
	MyFunc(MyFuncOpts{
		LastName: "Patel",
		Age:      23,
	})
	MyFunc(MyFuncOpts{
		FirstName: "Dave",
		LastName:  "John",
	})

}

//declaring function named div
func div(numerator int, denumerator int) int {
	if denumerator == 0 {
		return 0
	}
	return numerator / denumerator
}

//using struct to simulate named parameters
func MyFunc(opts MyFuncOpts) {
}

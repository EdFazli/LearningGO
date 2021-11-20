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

	//different ways to call variadic function named addTo
	fmt.Println(addTo(3))             // []
	fmt.Println(addTo(3, 2))          // [5]
	fmt.Println(addTo(3, 2, 4, 6, 8)) // [5 7 9 11]
	a := []int{4, 3}
	fmt.Println(addTo(3, a...))                    // [7 6]
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...)) // [4 5 6 7 8]

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

//variadic function named addTo
func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

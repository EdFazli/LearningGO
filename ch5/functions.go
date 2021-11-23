package main

import (
	"errors"
	"fmt"
	"os"
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

	//multiple return values
	results, remainder, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(results, remainder) // 2 1

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

//multiple return values
func divAndRemainder(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by error")
	}
	return numerator / denominator, numerator % denominator, nil
}

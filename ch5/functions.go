package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
)

//using struct to simulate named parameters
type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

//functions are values - with examples of calculator
type opFuncType func(int, int) int

//passing function as parameters
type Person struct {
	firstName string
	lastName  string
	age       int
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

	//functions are values - with examples of calculator
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}
	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression: ", expression)
			continue
		}
		p1, errorz := strconv.Atoi(expression[0])
		if errorz != nil {
			fmt.Println(errorz)
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator: ", op)
			continue
		}
		p2, errorz := strconv.Atoi(expression[2])
		if errorz != nil {
			fmt.Println(errorz)
			continue
		}
		result := opFunc(p1, p2)
		fmt.Println(result)
	}
	/*
		5
		-1
		6
		0
		unsupported operator:  %
		strconv.Atoi: parsing "two": invalid syntax
		invalid expression:  [5]
	*/

	//anonymous functions - useful for defer statement and launching goroutines
	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Println("printing", j, "from inside of anonymous function")
		}(i)
	}
	/*
		printing 0 from inside of anonymous function
		printing 1 from inside of anonymous function
		printing 2 from inside of anonymous function
		printing 3 from inside of anonymous function
		printing 4 from inside of anonymous function
	*/

	//passing function as parameters
	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobbert", 23},
		{"Fred", "Fredson", 18},
	}
	fmt.Println(people) // [{Pat Patterson 37} {Tracy Bobbert 23} {Fred Fredson 18}]

	sort.Slice(people, func(i int, j int) bool {
		return people[i].lastName < people[j].lastName
	})
	fmt.Println(people) // [{Tracy Bobbert 23} {Fred Fredson 18} {Pat Patterson 37}]

	sort.Slice(people, func(i int, j int) bool {
		return people[i].age < people[j].age
	})
	fmt.Println(people) // [{Fred Fredson 18} {Tracy Bobbert 23} {Pat Patterson 37}]

	//returning functions from functions
	twoBase := makeMult(2)
	threeBase := makeMult(3)
	for i := 0; i < 3; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}
	/*
		0 0
		2 3
		4 6
	*/

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

//functions are values - with examples of calculator
func add(i int, j int) int      { return i + j }
func sub(i int, j int) int      { return i - j }
func mul(i int, j int) int      { return i * j }
func division(i int, j int) int { return i / j }

var opMap = map[string]opFuncType{
	"+": add,
	"-": sub,
	"*": mul,
	"/": division,
}

//returning functions from functions
func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is functions examples")
	result := div(5, 2)
	fmt.Println(result)
}

//declaring function named div
func div(numerator int, denumerator int) int {
	if denumerator == 0 {
		return 0
	}
	return numerator / denumerator
}

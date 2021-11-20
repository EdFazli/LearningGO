package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("This is blocks, shadows and control structures examples")

	//shadowing variables - variable that has the same name as a variable in a containing block
	x := 10
	if x > 5 {
		fmt.Println(x) // 10
		x := 5
		fmt.Println(x) // 5
	}
	fmt.Println(x) // 10

	//shadowing with multiple assignment
	X := 10
	if X > 5 {
		X, Y := 5, 20
		fmt.Println(X, Y) // 5 20
	}
	fmt.Println(X) // 10

	//scoping variable to an if statement
	if x := rand.Intn(10); x == 0 {
		fmt.Println("That's too low")
	} else if x > 5 {
		fmt.Println("Thats too big: ", x)
	} else {
		fmt.Println("Thats a good number: ", x) // Thats a good number:  1
	}

}

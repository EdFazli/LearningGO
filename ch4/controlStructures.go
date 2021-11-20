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

	//complete for statement
	for i := 0; i < 10; i++ {
		fmt.Println(i) // 0 1 2 3 4 5 6 7 8 9
	}

	//condition-only for statement same as while statement in C
	i := 1
	for i < 100 {
		fmt.Println(i) // 1 2 4 8 16 32 64
		i *= 2
	}

	//infinite for statement - loops forever
	for {
		fmt.Println("This is infinite loop")
		if !(i == 4) { //CONDITION to exit from the loop
			break // to exit from the loop
		}
	}

	//using continue to make code clearer - replacing chain of if/else statement
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Naruto")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Boruto")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Sasuke")
			continue
		}
		fmt.Println(i) // prints i if the above condition not met
	}

	//for-range loop with a slice
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}
	/* prints position and value
	0 2
	1 4
	2 6
	3 8
	4 10
	5 12
	*/

	//ignoring the key in for-range loop
	for _, v := range evenVals {
		fmt.Println(v)
	}
	/* prints value only
	2
	4
	6
	8
	10
	12
	*/

	//ignoring the value in for-range loop
	uniqueNames := map[string]bool{
		"Fred":  true,
		"Raul":  true,
		"Wilma": true,
	}

	for k := range uniqueNames {
		fmt.Println(k)
	}
	/* prints key only
	Fred
	Raul
	Wilma
	*/

	//iterating over maps
	//map iteration order varies - security reason to prevent Hash DoS
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	for i := 0; i < 3; i++ {
		fmt.Println("Loop ", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}
	/*
		Loop  0
		b 2
		c 3
		a 1
		Loop  1
		a 1
		b 2
		c 3
		Loop  2
		a 1
		b 2
		c 3
	*/

	//iterating over strings
	samples := []string{"hello", "apple_n!"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}
	/*
		0 104 h
		1 101 e
		2 108 l
		3 108 l
		4 111 o

		0 97 a
		1 112 p
		2 112 p
		3 108 l
		4 101 e
		5 95 _
		6 110 n
		7 33 !
	*/

	//modifying the value doesnt modify the source
	oddVals := []int{1, 3, 5, 7, 9, 11}
	for _, v := range oddVals {
		v *= 2
	}
	fmt.Println(oddVals) // 1 3 5 7 9 11

	//labeling for statement - to escape nested loop
	examples := []string{"hello", "apple_n!"}
outer:
	for _, etc := range examples {
		for i, r := range etc {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue outer
			}
		}
		fmt.Println()
	}
	/*
		0 104 h
		1 101 e
		2 108 l
		0 97 a
		1 112 p
		2 112 p
		3 108 l
	*/

	//comparing for-range loop and complete for loop
	evenVals2 := []int{2, 4, 6, 8, 10}
	for i, v := range evenVals2 {
		if i == 0 {
			continue
		}
		if i == len(evenVals2)-2 {
			break
		}
		fmt.Println(i, v)
	}
	/*
		1 4
		2 6
	*/

	//recommended using standard loop when not iterating 1st element to the last element
	for i := 1; i < len(evenVals2)-2; i++ {
		fmt.Println(i, evenVals2[i])
	}
	/*
		1 4
		2 6
	*/

	//switch statement
	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length: ", wordLen)
		case 6, 7, 8, 9: // empty case means nothing happens
		default:
			fmt.Println(word, "is a long word!")
		}
	}
	/*
		a is a short word
		cow is a short word
		smile is exactly the right length:  5
		anthropologist is a long word!
	*/

	//switch with label
loop:
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, "is even")
		case i%3 == 0:
			fmt.Println(i, "is divisible by 3 not 2")
		case i%7 == 0:
			fmt.Println("exit the loop")
			break loop
		default:
			fmt.Println(i, "is boring")
		}
	}
	/*
		0 is even
		1 is boring
		2 is even
		3 is divisible by 3 not 2
		4 is even
		5 is boring
		6 is even
		exit the loop
	*/

	//
}

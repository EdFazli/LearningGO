package main

import "fmt"

func main() {
	fmt.Println("This is compostie types example")

	//Arrays declaration
	var x [3]int                   //declared but no value assigned
	var y = [4]int{15, 22, 35, 53} // declared with values

	//Sparse array - most elements are set to their zero value
	var z = [12]int{2, 3, 6: 55, 6, 2, 11: 100}

	//Using array literal to initialize array
	var j = [...]int{4, 43, 67, 86}

	//Compare array
	var k = [...]int{2, 44, 67, 534, 1}
	var i = [5]int{2, 44, 67, 534, 1}

	//2D array
	var p [2][3]int

	fmt.Println("X = ", x) // X =  [0 0 0]
	fmt.Println("Y = ", y) // Y =  [15 22 35 53]
	fmt.Println("Z = ", z) // Z =  [2 3 0 0 0 0 55 6 2 0 0 100]
	fmt.Println("J = ", j) // J =  [4 43 67 86]
	fmt.Println(k == i)    // prints true
	fmt.Println("P = ", p) // P =  [[0 0 0] [0 0 0]]
}

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

	fmt.Println("x = ", x) // X =  [0 0 0]
	fmt.Println("y = ", y) // Y =  [15 22 35 53]
	fmt.Println("z = ", z) // Z =  [2 3 0 0 0 0 55 6 2 0 0 100]
	fmt.Println("j = ", j) // J =  [4 43 67 86]
	fmt.Println(k == i)    // prints true
	fmt.Println("p = ", p) // P =  [[0 0 0] [0 0 0]]

	//len function - takes array and returns its length
	fmt.Println(len(z)) // 12

	//SLices declaration
	var O []int //returns nil
	var X = []int{5, 77, 86, 1, 22, 52}

	//Sparse slices - most elements are set to their zero value
	var Z = [12]int{12, 13, 6: 55, 16, 12, 11: 1100}

	//multidimensional slices
	var P [][]int

	fmt.Println("X = ", X) // X =  [5 77 86 1 22 52]
	fmt.Println("Z = ", Z) // Z =  [12 13 0 0 0 0 55 16 12 0 0 1100]
	fmt.Println("P = ", P) // P =  []

	//slice only can be compare with nil
	fmt.Println(X == nil) // prints false
	fmt.Println(O == nil) // prints true

	//Append - to grow slices
	var u []int
	fmt.Println("u = ", u) // u =  []
	u = append(u, 90)
	fmt.Println("u = ", u) // u =  [90]
	var U = []int{4, 23, 42, 55}
	fmt.Println("U = ", U) // U =  [4 23 42 55]
	U = append(U, 89)
	fmt.Println("U = ", U) // U =  [4 23 42 55 89]

	//expand source slice to individual value
	var r = []int{2, 4, 5, 6, 8}
	fmt.Println("r = ", r) // r =  [2 4 5 6 8]
	U = append(U, r...)
	fmt.Println("U = ", U) // U =  [4 23 42 55 89 2 4 5 6 8]

	//understanding capacity
	var C []int
	fmt.Println(C, len(C), cap(C)) // [] 0 0
	C = append(C, 10)
	fmt.Println(C, len(C), cap(C)) // [10] 1 1
	C = append(C, 20)
	fmt.Println(C, len(C), cap(C)) // [10 20] 2 2
	C = append(C, 30)
	fmt.Println(C, len(C), cap(C)) // [10 20 30] 3 4
	C = append(C, 40)
	fmt.Println(C, len(C), cap(C)) // [10 20 30 40] 4 4
	C = append(C, 50)
	fmt.Println(C, len(C), cap(C)) // [10 20 30 40 50] 5 8

	//make - to specify the type, length and optionally the capacity
	m := make([]int, 6)    // length and capacity = 6
	fmt.Println("m = ", m) // m =  [0 0 0 0 0 0]

	//specify initial length and capacity
	M := make([]int, 5, 8) // length = 5 and capacity = 8
	fmt.Println("M = ", M) // M =  [0 0 0 0 0]
	M = append(M, 44)
	fmt.Println("M = ", M) // M =  [0 0 0 0 0 44]

	//non-nil slice with length of 0
	n := make([]int, 0, 10) // length = 0 and capacity = 10
	fmt.Println("n = ", n)  // n =  []
	n = append(n, 1, 2, 3, 4, 5, 6)
	fmt.Println("n = ", n) // n =  [1 2 3 4 5 6]

	//slicing slices
	q := []int{2, 4, 6, 8}
	w := q[:2]
	a := q[1:]
	s := q[1:3]
	d := q[:]
	fmt.Println("q: ", q) // q:  [2 4 6 8]
	fmt.Println("w: ", w) // w:  [2 4]
	fmt.Println("a: ", a) // a:  [4 6 8]
	fmt.Println("s: ", s) // s:  [4 6]
	fmt.Println("d: ", d) // d:  [2 4 6 8]

	// examples of more confusing slices with the use of append - NOT RECOMMENDED!
	f := make([]int, 0, 5)
	f = append(f, 1, 2, 3, 4)
	g := f[:2]
	h := f[2:]
	fmt.Println(cap(f), cap(g), cap(h))
	g = append(g, 30, 40, 50)
	f = append(f, 60)
	h = append(h, 70)
	fmt.Println("f: ", f)
	fmt.Println("g: ", g)
	fmt.Println("h: ", h)

}

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

	//examples of more confusing/overwriting slices with the use of append - NOT RECOMMENDED!
	f := make([]int, 0, 5)
	f = append(f, 1, 2, 3, 4)
	g := f[:2]
	h := f[2:]
	fmt.Println(cap(f), cap(g), cap(h)) // 5 5 3
	g = append(g, 30, 40, 50)
	f = append(f, 60)
	h = append(h, 70)
	fmt.Println("f: ", f) // f:  [1 2 30 40 70]
	fmt.Println("g: ", g) // g:  [1 2 30 40 70]
	fmt.Println("h: ", h) // h:  [30 40 70]

	//full slice expression protect against append
	//g := f[:2:2] - has capacity of 2
	//h := f[2:4:4] - has capacity of 2

	//copy function
	F := []int{2, 4, 6, 8}
	G := make([]int, 3) // length of 3 will be copied from F
	num := copy(G, F)
	fmt.Println(G, num) // [2 4 6] 3

	//copy from middle of source slice
	Q := []int{1, 2, 3, 4}
	W := make([]int, 2)
	copy(W, Q[1:3])
	fmt.Println("W: ", W) // W:  [2 3]

	//converting rune to string
	var aa rune = 'x'
	var ss string = string(aa)
	var bb byte = 'y'
	var dd string = string(bb)

	fmt.Println("ss: ", ss) // ss:  x
	fmt.Println("dd: ", dd) // dd:  y

	//converting string to slices
	var qq string = "Hello World!"
	var qa []byte = []byte(qq)
	var qs []rune = []rune(qq)

	fmt.Println("qa: ", qa) // qa:  [72 101 108 108 111 32 87 111 114 108 100 33]
	fmt.Println("qs: ", qs) // qs:  [72 101 108 108 111 32 87 111 114 108 100 33]

	//declaring Maps
	var nilMap map[string]int
	fmt.Println("nilMap: ", nilMap) // nilMap:  map[]
	fmt.Println(nilMap == nil)      // prints true
	mapLiteral := map[string]int{}
	fmt.Println("mapLiteral: ", mapLiteral) // mapLiteral:  map[]
	teams := map[string][]string{
		"Developer": {"Mark", "Ralph", "George"},
		"Infra":     {"Ed", "Venc"},
	}
	fmt.Println("teams: ", teams) // teams:  map[Developer:[Mark Ralph George] Infra:[Ed Venc]]

	//declare Maps with default size
	ages := make(map[string]int, 10)
	ages["Simon"] = 24
	ages["Kira"] = 18

	fmt.Println(ages)          // map[Kira:18 Simon:24]
	fmt.Println(ages["Simon"]) // 24
	fmt.Println(ages["Kira"])  // 18

	//comma ok idiom - to check whether a key is in the map
	E := map[string]int{
		"sugar": 10,
		"salt":  15,
	}

	v, ok := E["sugar"]
	fmt.Println(v, ok) // 10 true
	V, ok := E["salt"]
	fmt.Println(V, ok) // 15 true
	vv, ok := E["cinnamon"]
	fmt.Println(vv, ok) // 0 false

	//deleting Maps
	ee := map[string]int{
		"hello":   2,
		"goodbye": 4,
	}
	delete(ee, "goodbye")

	//using map as a set
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, xxx := range vals {
		intSet[xxx] = true
	}
	fmt.Println(len(vals), len(intSet)) // 11 8
	fmt.Println(intSet[5])              // true
	fmt.Println(intSet[500])            // false
	if intSet[100] {
		fmt.Println("100 is in the set")
	}

	//struct - to group together any related data
	type person struct {
		name string
		age  int
		pet  string
	}

	//declare struct
	var fred person
	bob := person{
		"Bob",
		35,
		"snake",
	} // struct literal
	tasya := person{
		age:  44,
		pet:  "cat",
		name: "Tasya",
	}

	// a field in struct is accessed with dotted notation
	fmt.Println(tasya.pet) // cat
	fmt.Println(bob.age)   // 35
	fmt.Println(fred)      // { 0 }

	//anonymous struct - to translate external data to struct or vice versa
	var people struct {
		name  string
		age   int
		house string
	}
	people.name = "Kyle"
	people.age = 23
	people.house = "Condo"
	fmt.Println(people.name+",", people.age, "has a "+people.house) // Kyle, 23 has a Condo

	pet := struct {
		name string
		kind string
	}{
		name: "Garfield",
		kind: "cat",
	}
	fmt.Println(pet.name) // Garfield

}

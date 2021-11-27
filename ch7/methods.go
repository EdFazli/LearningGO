package main

import (
	"fmt"
	"math"
	"time"
)

type Person struct {
	firstName string
	lastName  string
	age       int
}

type Counter struct {
	total       int
	lastUpdated time.Time
}

type IntTree struct {
	val         int
	left, right *IntTree
}

//iota for enumerations
// type mailCategory int

// const (
// 	uncategorized mailCategory = iota
// 	personal
// 	spam
// 	social
// 	ads
// )

//embedding for composition
type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

// interfaces example - https://gobyexample.com/interfaces
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// interfaces example - https://gobyexample.com/interfaces
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

//Methods declaration
func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.firstName, p.lastName, p.age)
}

//pointer receivers and value receivers
func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

//methods for nil instances
func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

//>>>>>>>>>>>>>>>>>>>>>>>>MAIN FUNCTION<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<//

func main() {
	fmt.Println("This is types, methods and interfaces examples")

	//Methods invocation
	p := Person{
		firstName: "Dave",
		lastName:  "John",
		age:       24,
	}
	output := p.String()
	fmt.Println(output) // Dave John, age 24

	//pointer receivers and value receivers
	var c Counter
	fmt.Println(c.String()) // total: 0, last updated: 0001-01-01 00:00:00 +0000 UTC
	c.Increment()
	fmt.Println(c.String()) // total: 1, last updated: 2021-11-28 02:14:28.0163101 +0800 +08 m=+0.004450801

	//methods for nil instances
	var it *IntTree
	it = it.Insert(5)
	it = it.Insert(3)
	it = it.Insert(10)
	it = it.Insert(2)
	fmt.Println(it.Contains(2))  // true
	fmt.Println(it.Contains(12)) // false

	//embedding for composition
	m := Manager{
		Employee: Employee{
			Name: "Bob",
			ID:   "22341",
		},
		Reports: []Employee{},
	}

	fmt.Println(m.ID)            // 22341
	fmt.Println(m.Description()) // Bob (22341)

	// interfaces example - https://gobyexample.com/interfaces
	r := rect{width: 3, height: 4}
	C := circle{radius: 5}

	measure(r)
	measure(C)
	/*
		{3 4}
		12
		14
		{5}
		78.53981633974483
		31.41592653589793
	*/
}

//>>>>>>>>>>>>>>>>>>>>>>>>MAIN FUNCTION<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<//

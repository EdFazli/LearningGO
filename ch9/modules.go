package main

import (
	"fmt"

	print "github.com/EdFazli/LearningGO/tree/main/ch1/ch9/formatter"
	"github.com/EdFazli/LearningGO/tree/main/ch1/ch9/math"
)

//>>>>>>>>>>>>>>>>>>>>>>>>MAIN FUNCTION<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<//

func main() {
	fmt.Println("This is modules, packages and imports examples")

	//import example
	num := math.Double(2)
	output := print.Format(num)
	fmt.Println(output)
}

//>>>>>>>>>>>>>>>>>>>>>>>>MAIN FUNCTION<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<//

package ifstatements

import (
	"fmt"
	"math"
)

func QuadEquation() {
	a := 1.0
	b := 6.0
	c := 5.0
	fmt.Printf("a = %v\n", a)
	fmt.Printf("b = %v\n", b)
	fmt.Printf("c = %v\n", c)
	fmt.Println()

	if D := b*b - 4*a*c; D < 0 {
		fmt.Println("No solution")
	} else if D == 0 {
		fmt.Println("One solution: ", -b/(2*a))
	} else {
		fmt.Println("First solution: ", (-b - (math.Sqrt(D) / (2 * a))))
		fmt.Println("Second solution: ", (-b + math.Sqrt(D)/(2*a)))
	}
}

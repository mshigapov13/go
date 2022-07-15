package main

import (
	"fmt"

	"github.com/mshigapov13/go/basics/workflow/cycles"
)

func main() {
	val := 10
	fmt.Printf("Cheks is works like While by passing %d\n", val)
	fmt.Printf("Worked like While: - %t\n", cycles.IsWorksLikeWhile(val))
	fmt.Println()
	fmt.Println()

	val = 7
	fmt.Printf("Cheks is works like for on C lang by passing %d\n", val)
	cycles.LikeC(val)
}

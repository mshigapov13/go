package primitives

import "fmt"

var (
	found  = true
	Answer = 1
)

const AGGENDA = `GLOBALS - is initialised
found = true
Answer = 1

locals - is not initialised
Answer = 0
found = false
newAnwer - some pointer
`

func DisplayVariables() {
	fmt.Println(aggenda())
	fmt.Println("===========")
	globals()
	fmt.Println("===========")
	locals()
}

func aggenda() string {
	return AGGENDA
}

func globals() {
	fmt.Println("Showing globals:")
	fmt.Printf("found = %t\nAnswer = %d\n", found, Answer)
}

func locals() {
	var Answer int

	var (
		found     bool
		newAnswer = &Answer
	)
	fmt.Println("Showing locals:")
	fmt.Printf("found = %t\nAnswer = %d\nnewAnswer = %p\n", found, Answer, newAnswer)
}

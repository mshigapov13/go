package cycles

import "fmt"

func Forever() {
	for {
		_ = 2
	}
}

func IsWorksLikeWhile(val int) bool {
	i := 1
	for {
		if i*i >= val {
			break
		}
		i++
	}

	if i*i == val {
		return true
	}
	return false
}

func LikeC(val int) {
	for i := val; i > 0; i-- {
		fmt.Printf("%d bottles on table\n", i)
	}

	fmt.Printf("No bottles on table\n")
}

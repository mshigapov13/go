package owntypes

import "fmt"

type ErrorLevel int

type UserId int

func PrintCurrentErrorLevel(level ErrorLevel) {
	fmt.Println(level)
}

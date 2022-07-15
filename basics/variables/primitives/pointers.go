package primitives

import (
	"fmt"
	"unsafe"
)

func Adress() {
	answer := 2
	fmt.Printf("Value of value variable = %v\n", answer)

	answer_2 := new(int)
	fmt.Printf("Value of pointer variable which is inited by new() = %v\n", *answer_2)

	var answer_3 *int
	fmt.Printf("Value of pointer variable which is not inited = %v\n", *answer_3) //will cause panick
}

func DontDoIt() {
	vals := []int{10, 20, 30, 40, 50}
	start := unsafe.Pointer(&vals[0])
	size := unsafe.Sizeof(int(0))
	for i := 0; i < len(vals); i++ {
		item := *(*int)(unsafe.Pointer(uintptr(start) + size*uintptr(i)))
		fmt.Println(item)
	}
}

package main

import owntypes "github.com/mshigapov13/go/basics/variables/ownTypes"

func main() {

	owntypes.PrintCurrentErrorLevel(1)

	id := owntypes.UserId(3)
	owntypes.PrintCurrentErrorLevel(id) //incompatible parameter type

	_ = id
}

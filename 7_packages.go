package main

import (
	"fmt"
	"practice-go/util"
)

func main() {

	fmt.Println(util.Area(12.1, 12.1))

	// cannot refer to unexported name util.returnSomething
	// this is because its not capitalized.
	// util.returnSomething()

	// this will work.
	util.ReturnSomething()
}

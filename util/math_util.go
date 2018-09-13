// util is a separate package that reside inside practice-go source code
// any function defined inside util can be used by main package go files as long as
// we import util package first.

// A package will be initialised only once even if it is imported from multiple packages.
package util

import (
	"fmt"
	"log"
	"math"
)

var packageLevelVar = "PACKAGE_LEVEL"

// Every package can contain a init function.
// The init function should not have any return type and should not have any parameters.
func init() {
	fmt.Println("util package initialized")
	fmt.Println(packageLevelVar)
	log.Fatal("Just Saying")
	// log.Fatal has a oc.exit(1) at the end. So program should end.
}

// We capitalised the functions Area and Diagonal in the rectangle package.
// This has a special meaning in Go. Any variable or function which starts
// with a capital letter are exported names in go. Only exported functions
// and variables can be accessed from other packages.
// In this case we need to access Area and Diagonal functions from the main package.
// Hence they are capitalised.

func Area(len, wid float64) float64 {
	area := len * wid
	return area
}

func Diagonal(len, wid float64) float64 {
	diagonal := math.Sqrt((len * len) + (wid * wid))
	return diagonal
}

func returnSomething() string {
	return "I won't work"
}

func ReturnSomething() string {
	return "I will work"
}

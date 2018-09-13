// util is a separate package that reside inside practice-go source code
// any function defined inside util can be used by main package go files as long as
// we import util package first.
package util

import "math"

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

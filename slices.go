package main

import "fmt"

func main() {
	a := [5]int{76, 77, 78, 79, 80}
	var slice []int = a[0:4] //creates a slice from a[1] to a[3]
	fmt.Println(slice)

	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after", darr)

	// refer
	// https://github.com/golangbot/arraysandslices/blob/master/arrayslice.go
}

package main

import "fmt"

// A variadic function is a function that can accept variable number of arguments
func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

func main() {
	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 109)
	find(78, 38, 56, 98)

	// if the find function call has only one argument.
	// We have not passed any argument to the variadic nums ...int parameter.
	// This is perfectly legal and in this case nums is a nil slice with length and capacity 0.
	find(87)

	// There is a syntactic sugar which can be used to pass a slice to a variadic function.
	// You have to suffix the slice with ... If that is done, the slice is directly passed to
	// the function without a new slice being created.
	nums := []int{89, 90, 95}
	find(89, nums...)
}

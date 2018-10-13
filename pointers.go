package main

import "fmt"

func main() {
	b := 2551
	// *T is the type of the pointer variable which points to a value of type T.
	var a *int = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)

	var d *int
	fmt.Println("d is", d)
	d = &b
	fmt.Println("d after initialization is", d)

	// how to deference a pointer.
	fmt.Println("value of d is", *d)
	fmt.Println("Incrementing d")
	*d++
	fmt.Println("value of b is", b)
	fmt.Println("value of a is", *a)
}

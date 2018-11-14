package main

import "fmt"

func main() {
	// All elements in an array are automatically assigned the zero value of the array type.
	var a [3]int //int array with length 3
	fmt.Println(a)

	a[0] = 12 // array index starts at 0
	a[1] = 78
	a[2] = 50
	fmt.Println(a)

	// short hand declaration to create array
	b := []int{12, 78, 50}
	fmt.Println(b)

	c := [3]int{12}
	c[1] = 20
	fmt.Println(c)

	// ... makes the compiler determine the length
	unsized := [...]int{12, 78, 50}
	fmt.Println(unsized)

	//Arrays are value types
	//Arrays in Go are value types and not reference types.
	// This means that when they are assigned to a new variable,
	// a copy of the original array is assigned to the new variable.
	// If changes are made to the new variable, it will not be reflected in the original array.

	countryA := [...]string{"USA", "China", "India", "Germany", "France"}
	countryB := countryA // a copy of a is assigned to b
	countryB[0] = "Singapore"
	fmt.Println("Country A Pool is ", countryA)
	fmt.Println("Country B is ", countryB)

	//Similarly when arrays are passed to functions as parameters,
	// they are passed by value and the original array in unchanged.

	// range function.

	numArray := [...]float64{67.7, 89.8, 21, 78}
	sum := float64(0)
	for i, v := range numArray { //range returns both the index and value
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of a", sum)
}

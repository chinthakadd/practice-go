package main

import "fmt"

// Learn about use of const key word: Essentially, constant cannot be re-assigned with another value, immutable.
// The value of a constant should be known at compile time. Like Java, it MUST BE a compile time constant.
func main() {
	const a = 55 //allowed
	// a = 89 //reassignment not allowed

	// In the following assignment, a string is assigned to a named constant.
	// however the following assignment is an untyped assignment. Which means, essentially
	// the constant hello does not have  type.
	const hello = "Hello World"

	// But here, i am printing the type of hello constant
	// And I see that the type is printed as String.
	// Why? Because, Go is a Strongly Typed Language.
	// therefore all variables require a explicit type.
	// In the above scenario, hello constant is not assigned a type,
	// So in this statement, go assigns a type to the named constant
	// based on the assigned value. In this case, "Hello World" is a string,
	// therefore type of hello is a string.
	// NOTE: I am still not 100% clear on above yet.
	// coz the values are said to untyped in Go. So how does it know that this has to be
	// a string? OR string is the default type?
	// I figured it I think. It says that go figures out the type of the value
	// based on the syntax. That how it does assignments for int, float and complex types as well.
	fmt.Printf("Type of hello is %T", hello)

	// Strongly type the constant.
	// it is simple
	const helloWithType string = "Hello World"

	// In go, we can create alias for types
	// ex:
	type Literal string
	var name Literal = "name"
	fmt.Printf("\nType of name is %T", name)

	// However, Go does not allow re-assignments.
	// ex: the following will not work.
	// var customName string = ""
	// name = customName

	const aNumber = 5
	var intVar int = aNumber
	var int32Var int32 = aNumber
	var float64Var float64 = aNumber
	var complex64Var complex64 = aNumber
	fmt.Println("intVar",intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var",complex64Var)

}

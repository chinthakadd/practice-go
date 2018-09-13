package main

import (
	. "fmt"
	"math"
)

func main() {

	// https://golangbot.com/variables/
	// variable declaration
	var age int
	Println("age:", age)
	age = 29 //assignment
	Println("age:", age)
	age = 54 //assignment
	Println("age:", age)

	//declaring multiple variables
	// type can be omited.
	var width, height int = 100, 50
	Println("width:", width, "height:", height)

	var (
		id  = "1"
		ssn = "123-12-1111"
	)
	Println("id:", id, "ssn:", ssn)

	// short hand declaration
	subject, marks := "Maths", 98
	Println("subject:", subject, "marks:", marks)

	b, c := 40, 50 // b is already declared but c is new
	Println("b is", b, "c is", c)
	b, c = 80, 90 // assign new values to already declared variables b and c
	Println("changed b is", b, "c is", c)

	// illegal, declaration can only be done if at least variable on left side is new.
	// b, c := 70,12

	// functions
	var d = math.Min(12.1, 24.1)
	Println(d)

	// I am still unsure on which case to use for variables.
	const TestConst = "A"
	println(TestConst)
}

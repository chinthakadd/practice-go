package main

import "fmt"

func main() {
	// Default way of declaration.
	personSalary := make(map[string]int)
	personSalary["steve"] = 12000
	personSalary["jamie"] = 15000
	personSalary["mike"] = 9000

	fmt.Println("personSalary map contents:", personSalary)

	// It is also possible to initialize a map during declaration itself.
	personSalary2 := map[string]int{
		"oc":    12000,
		"jamie": 15000,
	}
	fmt.Println(personSalary2)

	//The map will return the zero value of the type of that element if the value
	// is not present.
	fmt.Println("Person Salary of Dude is:", personSalary["dude"])

	// value, ok := map[key]
	// The above is the syntax to find out whether a particular key is present in a map.
	// If ok is true, then the key is present and its value is present in the variable value, else the key is absent
	value, ok := personSalary["anotherDude"]
	fmt.Println(value, "Employee Exist:", ok)

	// for loop
	// One important fact is that the order of the retrieval of values from a map
	// when using for range is not guaranteed to be the same for each execution of the program.
	for key, value := range personSalary {
		fmt.Printf("personSalary[%s] = %d\n", key, value)
	}

	delete(personSalary, "jamie")
	fmt.Println("map after deletion", personSalary)
	fmt.Println("length is", len(personSalary))

	// Maps can't be compared using the == operator. The == can be only used to check if a map is nil.
}

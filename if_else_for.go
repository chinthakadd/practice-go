package main

import "fmt"

func main() {

	// num is initialised in the if statement.
	// One thing to be noted is that num is available only
	// for access from inside the if and else.
	// i.e. the scope of num is limited to the if else blocks.
	// If we try to access num from outside the if or else, the compiler will complain.
	num := 99
	if num <= 50 {
		fmt.Println("number is less than or equal to 50")
	} else if num >= 51 && num <= 100 {
		fmt.Println("number is between 51 and 100")
	} else {
		fmt.Println("number is greater than 100")
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}

	// short hand style
	i := 0
	for i <= 10 { //semicolons are ommitted and only condition is present
		fmt.Printf("%d ", i)
		i += 2
	}

	//n the above program no and i are declared and initialised to 10 and 1 respectively.
	// They are incremented by 1 at the end of each iteration.
	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 { //multiple initialisation and increment
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}

	//infinite loop
	//The syntax for creating an infinite loop is,
	//for {
	//	println("a")
	//}

}

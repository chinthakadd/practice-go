package main

import "fmt"

func main() {
	price := calculateBill(12, 1)
	fmt.Println(price)

	area, perimeter := rectProps(10.1, 20.5)
	fmt.Println("Area:", area, "Perimeter:", perimeter)

	// Blank Identifier
	// The rectProps function returns the area and perimeter of the rectangle.
	// What if we only need the area and want to discard the perimeter. This is where _ is of use.
	areaOnly, _ := rectProps(10.1, 20.5)
	fmt.Println("AreaOnly:", areaOnly)
}

// equals to `calculateBill(price int, no int) int`
func calculateBill(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}

// returning multiple variables from a function
func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}



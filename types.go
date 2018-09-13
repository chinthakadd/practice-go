package main

import (
	"fmt"

	// unsafe package should be used with care as
	// the code using it might have portability issues,
	// but for the purposes of this tutorial we can use it.
	"unsafe"
)

// The following are the basic types available in go
// bool
// Numeric Types

//	int8, int16, int32, int64, int
// 	we should generally be using int, which will represent int32 or int64 based on the platform.

//	uint8, uint16, uint32, uint64, uint
//	float32, float64

// The builtin function complex is used to construct a complex number with real and imaginary parts.
//	complex64, complex128

//	byte
//	rune
//string
func main() {
	var a int = 89
	b := 95
	fmt.Println("value of a is", a, "and b is", b)
	fmt.Println("type of a is %T, size of a is %d", a, unsafe.Sizeof(a))   //type and size of a
	fmt.Println("type of b is %T, size of b is %d", b, unsafe.Sizeof(b)) //type and size of b

	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("\nsum:", cadd)
	cmul := c1 * c2
	fmt.Println("product:", cmul)
}

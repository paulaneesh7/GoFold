package pgs2

import (
	"fmt"
)

func reference_types() {
	fmt.Println("Pointers and References")

	i, j := 42, 2701

	p := &i     		// point to i
	fmt.Println(*p)     // read i through the pointer
	*p = 21				// set i through the pointer
	fmt.Println(i)		// see the new value of i

	p = &j				// point to j
	*p = *p / 37		// divide j through the pointer
	fmt.Println(j)		// see the new value of j

}

func Basics2() {
	fmt.Println("More types: structs, slices, maps….")

	// Pointers and References
	reference_types()
	PointersGo()
}

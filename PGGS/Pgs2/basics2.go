package pgs2

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func reference_types() {
	fmt.Println("Pointers and References")

	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

}

func struct_usecase() {
	fmt.Println("Structs in Go")
	fmt.Println(Vertex{1, 2})

	v := Vertex{3,4}
	v.X = 10
	fmt.Println("Struct field after changing their value: ",v.X)

	p := &v
	p.X = 1e9
	fmt.Println("Struct field after changing their value using pointer: ",v)
}

func Basics2() {
	fmt.Println("More types: structs, slices, maps….")

	// Pointers and References
	reference_types()
	PointersGo()


	// structs
	struct_usecase()
}

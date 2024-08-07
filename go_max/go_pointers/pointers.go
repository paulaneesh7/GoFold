package go_pointers

import "fmt"

func PointerInGo() {

	fmt.Println("Pointers in Go")

	age := 32
	agePointer := &age // storing the address of age
	fmt.Println("Age: ", agePointer) // will show the address of age
	fmt.Println("Age: ", *agePointer) // will show the value stored at the address

	adultYears := getAdultYears(age)
	fmt.Println("Adult years: ", adultYears)

}

/*The age parameter is technially not the same age variable which we defined in the main func
It is a copy of it stored in a different place in memory.
*/
func getAdultYears(age int) int{
	return age - 18
}
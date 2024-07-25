package main

import "fmt"

/* Map is just like HashMap of Java */
func main() {
	// name <-> grade (mapping name with grade for every student)
	studentGrade := make(map[string]int)

	studentGrade["Aneesh"] = 100
	studentGrade["Swastik"] = 90
	studentGrade["Suvankar"] = 80
	studentGrade["Aniruddha"] = 100

	// retrieve the value of a key
	fmt.Println("Grade of Aneesh is: ", studentGrade["Aneesh"])
	fmt.Println("Grade of Swastik is: ", studentGrade["Swastik"])

	studentGrade["Swastik"] = 88
	fmt.Println("Updated grade of Swastik is: ", studentGrade["Swastik"])

	// delete a key
	delete(studentGrade, "Swastik")
	fmt.Println("Map after deletion: ", studentGrade)

	// check if a key is present
	CheckKeyPresent( studentGrade, "Aneesh")

	// iterate over the map
	for key, value := range studentGrade{
		fmt.Printf("Key is %s and grade is %d\n", key, value)
	}

	InitaliseMap()

}

func CheckKeyPresent(studentGrade map[string]int, name string){
	grade, isPresent := studentGrade[name]
	if isPresent{
		fmt.Println("Key is present and grade of ", name, " is: ", grade)
	}else {
		fmt.Println("Grade of ", name, " is not present")
	}
	
}

func InitaliseMap(){
	person := map[string]int{
		"Alice": 22,
		"Bob": 24,
		"Charlie": 29,
	}

	fmt.Println("Person map: ", person)
}
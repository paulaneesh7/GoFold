package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	var person1 Person // creating an instance of struct Person, so that we can access it's fields and properties later
	fmt.Println("Aneesh Person: ", person1)
	
	// 1st method to assign values to struct fields
	person1.FirstName = "Aneesh"
	person1.LastName = "Paul"
	person1.Age = 21
	fmt.Println("Aneesh Person: ", person1)

	// 2nd method to assign values to struct fields
	person2 := Person{
		FirstName: "Shiv",
		LastName:  "Kumar",
		Age: 22,
	}
	fmt.Println("Sh17v Person: ", person2)


	// 3rd method to assign values to struct fields
	person3 := new(Person)
	person3.FirstName = "Rahul"
	person3.LastName = "Dutta"
	person3.Age = 28
	fmt.Println("Rahul Person: ", person3)

	// multiple struct fields
	EmployeeDetails()
}

type Contact struct{
	Email string
	Phone string
}

type Address struct{
	House int
	Area string
	State string
}

type Employee struct{
	Person_Details Person
	Person_Contact Contact
	Person_Address Address
}

func EmployeeDetails(){
	var employee1 Employee
	employee1.Person_Details = Person{
		FirstName: "Ramesh",
		LastName: "Kumar",
		Age: 28,
	}

	employee1.Person_Contact = Contact{
		Email: "rameshkuamr.test@gmail.com",
		Phone: "9876543210",
	}

	employee1.Person_Address = Address{
		House: 123,
		Area: "Koramangala",
		State: "Karnataka",
	}

	fmt.Println("Employee1: ", employee1)
}
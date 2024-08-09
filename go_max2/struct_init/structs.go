package structinit

import (
	"fmt"
	"time"
)

/*type name_of_custom_type struct*/
type User struct {
	firstName string
	lastName string
	birthdate string
	createdAt time.Time
}

func StructInitial() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser User
	appUser = User{
		firstName: userFirstName,
		lastName: userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}

	// ... do something awesome with that gathered data!

	outputDetails(appUser)
}

/*Just like how we define params (name int) like that (u User) or (user User)*/
func outputDetails(u User){
	fmt.Println("First Name: ", u.firstName, "Last Name: ", u.lastName, "Birth Date: " , u.birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
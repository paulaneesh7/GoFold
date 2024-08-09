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

/* Attaching the function to the struct, so that's why (u User) at the begining */
func (u User) outputDetails(){
	fmt.Println("First Name: ", u.firstName, "Last Name: ", u.lastName, "Birth Date: " , u.birthdate)
}


/* We have used pointer here because otherwise it doesn't change the main-User struct's fields
It's creating a copy of its own and changing the fields of that copy */
func (u *User) clearUserName(){
	u.firstName = ""
	u.lastName = ""
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

	
	// ... do something awesome with that gathered data, for practse purpose we are using pointers here
	appUser.outputDetails()

	// changes the data in the struct
	appUser.clearUserName()
	appUser.outputDetails()
}



func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
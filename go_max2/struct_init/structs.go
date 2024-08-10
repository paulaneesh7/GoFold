package structinit

import (
	"fmt"
	"go_max2/user"
)




func StructInitial() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	/* as because we are returning a pointer from the constructor newUser, so "*User" here */
	var appUser *user.User
	appUser, err := user.NewUser(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	
	// ... do something awesome with that gathered data, for practse purpose we are using pointers here
	appUser.OutputDetails()

	// changes the data in the struct
	appUser.ClearUserName()
	appUser.OutputDetails()
}




func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
package user

import (
	"errors"
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

// Admin struct which embeds User struct, this is kinda like Inheritance - Java
type Admin struct {
	email string
	password string
	AdminUser User
}



/* Attaching the function to the struct, so that's why (u User) at the begining */
func (u User) OutputDetails(){
	fmt.Println("First Name: ", u.firstName, "Last Name: ", u.lastName, "Birth Date: " , u.birthdate)
}


/* We have used pointer here because otherwise it doesn't change the main-User struct's fields
It's creating a copy of its own and changing the fields of that copy */
func (u *User) ClearUserName(){
	u.firstName = ""
	u.lastName = ""
}


/* Struct constructor which resturns a pointer of the original, preventing from creating another */
func NewUser(firstName, lastName, birthdate string) (*User, error){

	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first Name, Last Name and Birthdate are required")
	}

	return &User{
		firstName: firstName,
		lastName: lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

// Admin constructor that contains nested-User struct
func NewAdmin(email, password string) Admin {
	return Admin{
		email: email,
		password: password,
		AdminUser: User{
			firstName: "Admin",
			lastName: "Admin",
			birthdate: "-------",
			createdAt: time.Now(),
		},
	}
}
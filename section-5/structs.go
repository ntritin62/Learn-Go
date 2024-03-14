package main

import (
	"fmt"

	"example.com/structs/user"
)



func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User

	appUser = &user.User{
		FirstName: "Tin",
	}
	if err != nil {
		fmt.Println(err)
		return
	}	

	// ... do something awesome with that gathered data!



	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}


func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

package main

import "fmt"

var MyVar = "some value"

func validatUserInput(remainingTicket uint) (bool, bool, string, uint) {
	var userName string
	var userTicket uint
	// ask the user for the input
	fmt.Printf("Enter you name:\n")
	fmt.Scan(&userName)
	isValidUseruserName := len(userName) >= 2
	fmt.Printf("How many ticket do you want:\n")
	fmt.Scan(&userTicket)
	isValidUseruserTicket := userTicket > 0 && userTicket < remainingTicket
	return isValidUseruserName, isValidUseruserTicket, userName, userTicket

}

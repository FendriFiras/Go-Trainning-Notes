package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// go can apply the data type based on the Value
//var eventName = "Go conference"
var eventName = "Go conference" // syntac sugar
const conferenceTicket = 50

// uint type will accept only positive numbers
var remainingTicket uint = 50

//var bookings = [50]string{}
var bookings = make([]user, 0) //declaring it as a slice
// bookings := []string{}

//creating a struct
type user struct {
	name         string
	ticketNumber uint
}

var wg = sync.WaitGroup{}

func main() {

	//fmt.Printf("welcome to our %v booking application \n", eventName)

	//calling the funciton
	greetUsers()

	// fmt.Printf("We have total of %v Tickets\n", conferenceTicket)
	// fmt.Printf("Book your ticket now (Only %v available)\n", remainingTicket)

	fmt.Printf("the type of eventName is: %T\n", eventName)

	isValidUseruserName, isValidUseruserTicket, userName, userTicket := validatUserInput(remainingTicket)
	if isValidUseruserName && isValidUseruserTicket {
		if remainingTicket == 0 {
			fmt.Printf("Our Conference is booked our. come back next year\n")
		} else {
			remainingTicket = remainingTicket - userTicket

			// creat a map for a user
			// var userData = make(map[string]string)
			// userData["name"] = userName
			// userData["numberTickets"] = strconv.FormatUint(uint64(userTicket), 10)

			var userData = user{
				name:         userName,
				ticketNumber: userTicket,
			}
			bookings = append(bookings, userData)

			// we just made our application concurrent
			wg.Add(1)
			go sendTicket(userName, userTicket)

			info(userName, userTicket, remainingTicket, bookings)

			booking := booclFor(bookings)
			fmt.Printf("The type of the array: %v \n", booking)

		}
	} else if userTicket == remainingTicket {
		fmt.Printf("You get all remainning tickets\n")
	} else {
		if !isValidUseruserName {
			fmt.Printf("your username input data is invalid,try again\n")
		}
		if !isValidUseruserTicket {
			fmt.Printf("your Tickets numbr input data is invalid,try again\n")
		}

	}

	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v we have a total of %v and just %v remaining\n", eventName, conferenceTicket, remainingTicket)
}

func booclFor(bookings []user) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, strconv.FormatUint(uint64(booking.ticketNumber), 10))
		// fmt.Printf("the index is %v", index)
	}
	fmt.Printf(" the list of byed tickets%v\n", firstNames)
	return firstNames
}

func info(userName string, userTicket uint, remainingTicket uint, bookings []user) {
	fmt.Printf("user %v booked %v tickets \n", userName, userTicket)
	fmt.Printf("Remaining tickets: %v", remainingTicket)
	fmt.Printf("The whole array: %v \n", bookings)
	fmt.Printf("The lenth of the array: %v\n", len(bookings))

}

func sendTicket(userName string, userTicket uint) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v this many tickets for %v\n", userTicket, userName)
	fmt.Printf("sending ticket to u %v\n", ticket)

	wg.Done()
}

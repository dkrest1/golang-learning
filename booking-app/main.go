package main

import (
	"booking-app/helpers"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50 
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string;
	lastName string;
	email string;
	userTickets uint
}

var wg = sync.WaitGroup{}


func main () {
	
	greeUser()

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isvalidTicketNumber := 
		helpers.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if   isValidName && isValidEmail && isvalidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames :=getFirstNames()
			fmt.Printf("These are the names %v of the poople that have booked for the conference \n", firstNames)

			noTicketRemaining  := remainingTickets == 0
			if  noTicketRemaining {
				fmt.Println("Sorry our Conference is booked out, come back next year.")
				// break
			}
		}else {

			if !isValidName {
				fmt.Println("Please enter a valid name")
			}
			if !isValidEmail {
				fmt.Println("the email you entered is invlid")
			}

			if !isvalidTicketNumber {

				fmt.Println("You have entered an invlid ticker number")
			}
		}
			
	wg.Wait()
}

func greeUser() {	
	fmt.Printf("You are welcome to %v \n", conferenceName)
	fmt.Printf("We have total of %v ticket and %v are still available. \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend now")
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {

	var firstName string

	var lastName string

	var email string

	var userTickets uint

	fmt.Println("Enter your firstName: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your lastName: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter how many tickets you want to get: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create a ticket
	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		userTickets: userTickets,
	} 

	bookings = append(bookings, userData)

	fmt.Printf("List of booking %v \n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets, you will receive a comfirmation email at %v \n", firstName, lastName, userTickets, email)

	fmt.Printf("%v remaining tickets left for %v \n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#############################")
	fmt.Printf("Sending Ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#############################")
	wg.Done()
}
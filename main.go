package main

import (
	// "booking-app/helper"
	F "fmt"
	"sync"
	"time"
)

var conferenceName string = "Go Conference"

const conferenceTickets = 50

var conferenceRemainingTickets uint = 50

// Create a list of type struct with empty list
var bookingsList = make([]UserData, 0)

// Assign variable names to and data type in struct
type UserData struct {
	firstName       string
	lastName        string
	emailAddress    string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	// call gretings function
	greetUsers()

	// Get user detials
	firstName, lastName, emailAddress, userTickets := getUserInfo()

	// Vaidate user details
	isValidName, isValidEmail, isValidTickets := ValidateUserInput(firstName, lastName, emailAddress, userTickets, conferenceRemainingTickets)
	if isValidName && isValidEmail && isValidTickets {

		// book user ticket
		conferenceRemainingTickets := bookTicket(firstName, lastName, emailAddress, userTickets)

		wg.Add(1)

		// Send email to user with ticket
		go SendTicket(userTickets, firstName, lastName, emailAddress)

		// Call function to print first name
		F.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email shortly at %v.\n", firstName, lastName, userTickets, emailAddress)
		F.Printf("%v tickets remaining for %v\n", conferenceRemainingTickets, conferenceName)

		// Retrieve user first name
		firstNames := getFirstName()
		F.Printf("The first names of booking are %v\n", firstNames)

		// confirm amount of tickets
		noTicketRemaining := conferenceRemainingTickets == 0
		if noTicketRemaining {
			F.Printf("Sorry, we are sold out. Please try again later.\n")
			//break
		}
	} else {
		if !isValidEmail {
			F.Println("Please enter a valid email address")
		}
		if !isValidTickets {
			F.Println("Please enter a valid amount of tickets")
		}
		if !isValidName {
			F.Println("Please enter a valid name")
		}

	}

	// Wait for goroutines to finish before exiting the program
	wg.Wait()

}

// Greet user function
func greetUsers() {
	// greet users when they get to the screen
	F.Printf("Welcome to %s booking application\n", conferenceName)
	F.Printf("We have total of %d tickets and %d are still available.\n", conferenceTickets, conferenceRemainingTickets)
	F.Printf("Get your tickets here to attend the %v\n", conferenceName)
}

// Get user info function
func getUserInfo() (string, string, string, uint) {
	var firstName string
	var lastName string
	var emailAddress string
	var userTickets uint

	// ask user for name
	F.Println("Please enter your first name:")
	F.Scanln(&firstName)

	F.Println("Please enter your last name:")
	F.Scanln(&lastName)

	F.Println("Please enter your email address:")
	F.Scanln(&emailAddress)

	F.Println("Please enter amount of tickets you would like to get:")
	F.Scanln(&userTickets)

	return firstName, lastName, emailAddress, userTickets
}

// Retrieve firstname from the list function
func getFirstName() []string {
	firstNames := []string{}

	for _, booking := range bookingsList {

		// call user first name
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

// book user tickets function
func bookTicket(firstName, lastName, emailAddress string, userTickets uint) uint {
	conferenceRemainingTickets -= userTickets

	// Create a type struct for a user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		emailAddress:    emailAddress,
		numberOfTickets: userTickets,
	}

	bookingsList = append(bookingsList, userData)
	F.Println(bookingsList)

	return conferenceRemainingTickets
}

// send user confirmation email and tickets function
func SendTicket(userTickets uint, firstName, lastName, emailAddress string) {

	// Sleep for 10 seconds
	time.Sleep(time.Second * 10)

	tickets := F.Sprintf("%v tickets for %v %v \n", userTickets, firstName, lastName)
	F.Println("####################################")
	F.Printf("sending ticket:\n %v \nto email address:\n %v\n", tickets, emailAddress)
	F.Println("####################################")

	// Decrement wait group to indicate that the goroutine is done
	wg.Done()

}

// Go task
// write a go program that takes a link and use go routines to find broken links

package main

import (
	F "fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets = 50
	var conferenceRemainingTickets uint = 50
	var bookingsList []string

	// call gretings function

	for {
		// Greet user
		greetUsers(conferenceName, conferenceTickets, conferenceRemainingTickets)

		// Get user detials
		firstName, lastName, emailAddress, userTickets := getUserInfo()

		// Vaidate user details
		isValidName, isValidEmail, isValidTickets := validateUserInput(firstName, lastName, emailAddress, conferenceRemainingTickets, userTickets)
		if isValidName && isValidEmail && isValidTickets {

			// book user ticket
			bookingsList, conferenceRemainingTickets := bookTicket(bookingsList, firstName, lastName, conferenceName, emailAddress, conferenceRemainingTickets, userTickets)

			// Call function to print first name
			F.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email shortly at %v.\n", firstName, lastName, userTickets, emailAddress)
			F.Printf("%v tickets remaining for %v\n", conferenceRemainingTickets, conferenceName)

			// Retrieve user first name
			firstNames := getFirstName(bookingsList)
			F.Printf("The first names of booking are %v\n", firstNames)

			// confirm amount of tickets
			noTicketRemaining := conferenceRemainingTickets == 0
			if noTicketRemaining {
				F.Printf("Sorry, we are sold out. Please try again later.\n")
				break
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

	}

}

// Greet users

func greetUsers(conferenceName string, conferenceTickets int, conferenceRemainingTickets uint) {
	// greet users when they get to the screen
	F.Printf("Welcome to %s booking application\n", conferenceName)
	F.Printf("We have total of %d tickets and %d are still available.\n", conferenceTickets, conferenceRemainingTickets)
	F.Printf("Get your tickets here to attend the %v\n", conferenceName)
}

// Get user info
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

// Retrieve firstname from the list
func getFirstName(bookingsList []string) []string {
	firstNames := []string{}

	for _, booking := range bookingsList {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

// Validate user info
func validateUserInput(firstName, lastName, emailAddress string, conferenceRemainingTickets, userTickets uint) (bool, bool, bool) {

	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(emailAddress, "@") && strings.Contains(emailAddress, ".")
	isValidTickets := userTickets > 0 && userTickets <= conferenceRemainingTickets

	return isValidName, isValidEmail, isValidTickets
}

// book user tickets
func bookTicket(bookingsList []string, firstName, lastName, conferenceName, emailAddress string, conferenceRemainingTickets uint, userTickets uint) ([]string, uint) {
	bookingsList = append(bookingsList, firstName+" "+lastName)
	conferenceRemainingTickets -= userTickets

	return bookingsList, conferenceRemainingTickets
}

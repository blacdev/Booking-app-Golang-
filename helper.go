/*
Here, the validateUserInput function was pulled out
to create a file that ca be re used all over the package
*/
package main

import (
	"strings"
)

// Validate user info
func ValidateUserInput(firstName, lastName, emailAddress string, userTickets, conferenceRemainingTickets uint) (bool, bool, bool) {

	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(emailAddress, "@") && strings.Contains(emailAddress, ".")
	isValidTickets := userTickets > 0 && userTickets <= conferenceRemainingTickets

	return isValidName, isValidEmail, isValidTickets
}

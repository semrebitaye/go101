package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTicket uint, remainingTicket uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTicket > 0 && remainingTicket >= userTicket

	return isValidName, isValidEmail, isValidTicketNumber
}

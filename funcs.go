package main

import "strings"

func ValidateUserInput(userName string, userEmail string, userTickets uint, city string) (bool, bool, bool, bool) {
	isNameValid := len(userName) >= 2
	isEmailValid := strings.Contains(userEmail, "@") && strings.Contains(userEmail, ".") && len(userEmail) >= 3
	isUserTicketsValid := userTickets > 0 && userTickets <= remainingtick
	isValidCity := city == "Singapore" || city == "London"
	return isNameValid, isEmailValid, isUserTicketsValid, isValidCity
}

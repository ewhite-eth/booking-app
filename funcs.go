package main

import "strings"

func ValidateUserInput(userName string, userEmail string, userTickets uint, city string) (bool, bool, bool, bool) {
	isNameValid := len(userName) >= 2
	// Basic email validation: check for '@' and '.' with proper positioning
	atIndex := strings.Index(userEmail, "@")
	dotIndex := -1
	if atIndex > 0 {
		if dotAfterAt := strings.Index(userEmail[atIndex:], "."); dotAfterAt != -1 {
			dotIndex = dotAfterAt + atIndex
		}
	}
	isEmailValid := atIndex > 0 && dotIndex > atIndex+1 && dotIndex < len(userEmail)-1
	isUserTicketsValid := userTickets > 0 && userTickets <= remainingtick
	isValidCity := city == "Singapore" || city == "London"
	return isNameValid, isEmailValid, isUserTicketsValid, isValidCity
}

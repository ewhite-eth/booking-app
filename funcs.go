package main

func ValidateUserInput(userName string, userTickets uint, city string) (bool, bool, bool) {
	isNameValid := len(userName) >= 2
	isUserTicketsValid := userTickets > 0 && userTickets <= remainingtick
	isValidCity := city == "Singapore" || city == "London"
	return isNameValid, isUserTicketsValid, isValidCity
}

// stvor creates and returns a new UserData struct
func stvor(userName string, userTickets uint, city string) UserData {
	return UserData{
		userName:    userName,
		userTickets: userTickets,
		city:        city,
	}
}

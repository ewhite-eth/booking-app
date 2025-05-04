package main


func ValidateUserInput(userName string, userTickets uint, city string) (bool, bool, bool) {
    isNameValid := len(userName) >= 2
    isUserTicketsValid := userTickets > 0 && userTickets <= remainingtick
    isValidCity := city == "Singapore" || city == "London"
    return isNameValid, isUserTicketsValid, isValidCity
}
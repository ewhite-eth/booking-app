package main

import (
	"fmt"
	"sync"
	"time"
)

// Constants
const tickets uint = 50
const ticketSendingDelay = 10 * time.Second

var confName = "Go Conference"
var remainingtick uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	userName    string
	userTickets uint
	userEmail   string
	city        string
}

func main() {
	greetUser()

	var wg = sync.WaitGroup{}

	for remainingtick > 0 {
		userName, userEmail, userTickets, city := getUserInput()
		isNameValid, isEmailValid, isUserTicketsValid, isValidCity := ValidateUserInput(userName, userEmail, userTickets, city)

		if isNameValid && isEmailValid && isUserTicketsValid && isValidCity {
			remainingtick = booking(remainingtick, userTickets, userName, userEmail, city)

			wg.Add(1)
			go sendTicket(userTickets, userName, userEmail, &wg)

			fmt.Printf("These are all bookings: %v\n", bookings)

			if remainingtick == 0 {
				fmt.Println("Conference is booked out")
				break
			}
		} else {
			if !isNameValid {
				fmt.Println("Invalid name. Must be at least 2 characters.")
			}
			if !isEmailValid {
				fmt.Println("Invalid email. Must contain '@' and '.' symbols.")
			}
			if !isValidCity {
				fmt.Println("Invalid city. Choose Singapore or London.")
			}
			if !isUserTicketsValid {
				fmt.Printf("Invalid ticket number. Only %v tickets left.\n", remainingtick)
			}
			fmt.Println("Please try again.")
		}
	}

	wg.Wait()
}

func greetUser() {
	fmt.Printf("\n----------------------Welcome to %v booking application----------------------\n\n", confName)
	fmt.Printf("We have a total of %v tickets and %v of them are available\n", tickets, remainingtick)
	fmt.Println("Get your tickets to attend")
}

func getUserInput() (string, string, uint, string) {
	var userName string
	var userEmail string
	var userTickets uint
	var city string

	fmt.Print("Enter your username: ")
	fmt.Scan(&userName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&userEmail)

	fmt.Print("How many tickets: ")
	fmt.Scan(&userTickets)

	fmt.Print("Enter your city (Singapore or London): ")
	fmt.Scan(&city)

	return userName, userEmail, userTickets, city
}

func booking(remainingtick uint, userTickets uint, userName string, userEmail string, city string) uint {
	remainingtick -= userTickets
	//create a map to store the user name and tickets
	var userData = UserData{
		userName:    userName,
		userTickets: userTickets,
		userEmail:   userEmail,
		city:        city,
	}

	bookings = append(bookings, userData)
	if userTickets == 1 {
		fmt.Printf("User %v booked 1 ticket.\n", userName)
	} else {
		fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
	}
	fmt.Printf("%v tickets are remaining\n", remainingtick)
	return remainingtick
}

func sendTicket(userTickets uint, userName string, userEmail string, wg *sync.WaitGroup) {
	time.Sleep(ticketSendingDelay)
	var ticket = fmt.Sprintf("%v tickets for %v", userTickets, userName)
	fmt.Println("##################")
	fmt.Printf("Sending ticket:\n%v\nto email address %v\n", ticket, userEmail)
	fmt.Println("##################")
	wg.Done()
}

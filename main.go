package main

import (
	"fmt"
)

// Constants
const tickets uint = 50
var confName = "Go Conference"
var remainingtick uint = 50
var bookings= make([] UserData,0)

type UserData struct  {
    userName string
    userTickets uint
    city string
}

func main() {
    greetUser()

    for remainingtick > 0 {
        userName, userTickets, city := getUserInput()
        isNameValid, isUserTicketsValid, isValidCity := ValidateUserInput(userName, userTickets, city)

        if isNameValid && isUserTicketsValid && isValidCity {
            remainingtick = booking(remainingtick, userTickets, userName, city)

            fmt.Printf("These are all bookings: %v\n", bookings)

            if remainingtick == 0 {
                fmt.Println("Conference is booked out")
                break
            }
        } else {
            if !isNameValid {
                fmt.Println("Invalid name. Must be at least 2 characters.")
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
}

func greetUser() {
    fmt.Printf("\n----------------------Welcome to %v booking application----------------------\n\n", confName)
    fmt.Printf("We have a total of %v tickets and %v of them are available\n", tickets, remainingtick)
    fmt.Println("Get your tickets to attend")
}

func getUserInput() (string, uint, string) {
    var userName string
    var userTickets uint
    var city string

    fmt.Print("Enter your username: ")
    fmt.Scan(&userName)

    fmt.Print("How many tickets: ")
    fmt.Scan(&userTickets)

    fmt.Print("Enter your city (Singapore or London): ")
    fmt.Scan(&city)

    return userName, userTickets, city
}

func booking(remainingtick uint, userTickets uint, userName string, city string) uint {
    remainingtick -= userTickets
    //create a map to store the user name and tickets
    var userData=UserData {
        userName: userName,
        userTickets: userTickets,
        city: city,
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
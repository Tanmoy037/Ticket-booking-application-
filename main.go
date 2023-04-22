package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	var conferenceTicktes uint = 50
	var remainingTicktes uint = 50
	bookings := []string{}

	fmt.Printf("Welcome %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTicktes ,"tickets and", remainingTicktes, "are still available")
	fmt.Println("Get your ticktes here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTicket int
		fmt.Println("Enter first your name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email: ")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTicket)


		remainingTicktes = remainingTicktes - uint(userTicket)
		bookings = append(bookings, firstName + " " + lastName)

		if userTicket > int(remainingTicktes) {
			fmt.Printf("We only have %v tickets remaining. So you can't book %v tickets. \n", remainingTicktes, userTicket)
			continue
			
		}


		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTicket, email)
		fmt.Printf("%v tickets remaining for %v \n", remainingTicktes, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])

		}


		fmt.Printf("These first names of bookings are: %v\n", firstName )


		if remainingTicktes == 0 { 
			
			fmt.Println("Our conference is booked out. Come back next year.")
			
			break
		}
	}

}
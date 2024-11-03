package main

import "fmt"

func main() {
	// We define variables and Go infers the types when we assign values
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	// Using println
	//fmt.Println("Welcome to", conferenceName, "booking application")
	//fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	//fmt.Println("Get your tickets here to attend the conference")

	// Using printf (formatted data)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend the conference\n")

	var userName string
	var userTickets int
	// ask the user for their input
	userName = "Tom"
	userTickets = 2
	// Print the user input
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
	// Print out the types of variables
	fmt.Printf("conferenceTickets is of type %T, remaining Tickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

}

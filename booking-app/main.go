package main

import "fmt"

func checkError(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err)
		return
	}
}

func main() {
	// We define variables and Go infers the types when we assign values
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	// Using println
	//fmt.Println("Welcome to", conferenceName, "booking application")
	//fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	//fmt.Println("Get your tickets here to attend the conference")

	// Using printf (formatted data)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend the conference\n")

	// Print out the types of variables
	fmt.Printf("conferenceTickets is of type %T, remaining Tickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	var userName string
	var userTickets uint
	var userEmail string

	// fmt.Scan is used for input from the user, similar to readline in node.js

	// Get the user input
	fmt.Print("\n\nEnter your first name: ")
	_, userNameErr := fmt.Scan(&userName)
	checkError(userNameErr, "Error while getting userName:")

	fmt.Print("Enter the number of tickets: ")
	_, userTicketsErr := fmt.Scan(&userTickets)
	checkError(userTicketsErr, "Error while getting userTickets:")

	fmt.Print("Enter your email address: ")
	_, userEmailErr := fmt.Scan(&userEmail)
	checkError(userEmailErr, "Error while getting userEmail:")

	// Reduce the number of remaining tickets
	remainingTickets = remainingTickets - userTickets

	// Print the user input
	fmt.Printf("\n\nUser %v booked %v tickets.\n", userName, userTickets)

	// Thank the user
	fmt.Printf("\n\nThank you %v for booking %v tickets for the %v conference. You will receive a confirmation email at %v.\n", userName, userTickets, conferenceName, userEmail)

	// Print remaining tickets
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	// Lists, arrays, and slices
	// Data structures to store collections of elements in a single variable
	// Arrays are fixed size, slices are dynamic

	// Arrays
	//var bookings = [50]string{"Nana", "Nicole", "Peter"}
	// We cannot mix types, it can only be of one type of data
	// We can add elements up to the maximum size of 50
	// We can start with an empty list also with
	// var bookings = [50]string or var bookings [50]string{}

	var bookings [50]string
	bookings[0] = userName

}

package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference" // this is an other way to declare a variable
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	//var bookingsarray = [50]string{}
	//this is an array

	//var bookingsSlice = []string{}
	//this is a slice it can also written like this : bookingsSlice := []string{}

	bookings := []string{}

	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("we have total of %v available and %v Remaining \n", conferenceTickets, remainingTickets)
	fmt.Println("get your ticket here to attend")

	// for { //this is an infinite loop
	// for true { //this is an infinite loop too

	for remainingTickets > 0 && len(bookings) > 50 { // conditional loop
		var firstName string
		var lastName string
		var email string

		var userTickets uint

		// ask user for info
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tckets: ")
		fmt.Scan(&userTickets)

		if userTickets < remainingTickets {
			// Array
			// bookingsarray[0] = firstName + " " + lastName
			// fmt.Printf("the whole array: %v \n", bookingsarray)
			// fmt.Printf("the first valye of array: %v \n", bookingsarray[0])
			// fmt.Printf("the array type: %T \n", bookingsarray)
			// fmt.Printf("the array length: %v \n", len(bookingsarray))

			// Slice
			// bookingsSlice = append(bookingsSlice, firstName+" "+lastName)
			// fmt.Printf("the whole slice: %v \n", bookingsSlice)
			// fmt.Printf("the first valye of slice: %v \n", bookingsSlice[0])
			// fmt.Printf("the slice type: %T \n", bookingsSlice)
			// fmt.Printf("the slice length: %v \n", len(bookingsSlice))

			bookings = append(bookings, firstName+" "+lastName+"|")
			fmt.Printf("the whole slice: %v \n", bookings)
			fmt.Printf("the first valye of slice: %v \n", bookings[0]) // %v returns the value
			fmt.Printf("the slice type: %T \n", bookings)              // %T returns the Type
			fmt.Printf("the slice length: %v \n", len(bookings))

			remainingTickets = remainingTickets - userTickets

			fmt.Printf("user %v %v with the email address of %v booked %v tickets. \n", firstName, lastName, email, userTickets)
			fmt.Printf("Remaining tickets %v for %v. \n", remainingTickets, conferenceName)

			firstNames := []string{} //slice

			for _, booking := range bookings { //loop slice elementes
				var names = strings.Fields(booking)           // seperate each element(booking) of bookings into a slice
				firstNames = append(firstNames, names[0]+"|") //names[0] returns the first world of each booking and than we append it into an other slice
			}
			fmt.Printf("the first names of All bookings:  %v. \n", firstNames)

			fmt.Printf("All bookings:  %v. \n", bookings)

			if remainingTickets == 0 {
				fmt.Printf("our %v is booked out. come back next year. ", conferenceName)
				break
			}
		} else if userTickets == remainingTickets {
			fmt.Printf("you are bying all the available tickets we have ! ")
		} else {
			fmt.Printf("we only have %v avaialble, you cant book %v. \n", remainingTickets, userTickets)
		}

	}

}

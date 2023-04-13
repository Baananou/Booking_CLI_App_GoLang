package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}


func main() {

	greetUsers()

	//conferenceName := "Go Conference" // this is an other way to declare a variable
	//const conferenceTickets int = 50
	//var remainingTickets uint = 50
	//var bookingsarray = [50]string{}
	//this is an array

	//var bookingsSlice = []string{}
	//this is a slice it can also written like this : bookingsSlice := []string{}

	// bookings := []string{}

	// fmt.Printf("Welcome to %v booking application \n", conferenceName)
	// fmt.Printf("we have total of %v available and %v Remaining \n", conferenceTickets, remainingTickets)
	// fmt.Println("get your ticket here to attend")

// for {

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			// end program
			fmt.Println("Our conference is booked out. Come back next year.")
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("first name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("email address you entered doesn't contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("number of tickets you entered is invalid")
		}
	}

	//}

	wg.Wait()

	// for { //this is an infinite loop
	// for true { //this is an infinite loop too

	// for remainingTickets > 0 && len(bookings) > 50 { // conditional loop
	// 	var firstName string
	// 	var lastName string
	// 	var email string

	// 	var userTickets uint

	// 	// ask user for info
	// 	fmt.Println("Enter your first name: ")
	// 	fmt.Scan(&firstName)

	// 	fmt.Println("Enter your last name: ")
	// 	fmt.Scan(&lastName)

	// 	fmt.Println("Enter your email: ")
	// 	fmt.Scan(&email)

	// 	fmt.Println("Enter number of tckets: ")
	// 	fmt.Scan(&userTickets)

	// 	if userTickets < remainingTickets {
	// 		// Array
	// 		// bookingsarray[0] = firstName + " " + lastName
	// 		// fmt.Printf("the whole array: %v \n", bookingsarray)
	// 		// fmt.Printf("the first valye of array: %v \n", bookingsarray[0])
	// 		// fmt.Printf("the array type: %T \n", bookingsarray)
	// 		// fmt.Printf("the array length: %v \n", len(bookingsarray))

	// 		// Slice
	// 		// bookingsSlice = append(bookingsSlice, firstName+" "+lastName)
	// 		// fmt.Printf("the whole slice: %v \n", bookingsSlice)
	// 		// fmt.Printf("the first valye of slice: %v \n", bookingsSlice[0])
	// 		// fmt.Printf("the slice type: %T \n", bookingsSlice)
	// 		// fmt.Printf("the slice length: %v \n", len(bookingsSlice))

	// 		bookings = append(bookings, firstName+" "+lastName+"|")
	// 		fmt.Printf("the whole slice: %v \n", bookings)
	// 		fmt.Printf("the first valye of slice: %v \n", bookings[0]) // %v returns the value
	// 		fmt.Printf("the slice type: %T \n", bookings)              // %T returns the Type
	// 		fmt.Printf("the slice length: %v \n", len(bookings))

	// 		remainingTickets = remainingTickets - userTickets

	// 		fmt.Printf("user %v %v with the email address of %v booked %v tickets. \n", firstName, lastName, email, userTickets)
	// 		fmt.Printf("Remaining tickets %v for %v. \n", remainingTickets, conferenceName)

	// 		firstNames := []string{} //slice

	// 		for _, booking := range bookings { //loop slice elementes
	// 			var names = strings.Fields(booking)           // seperate each element(booking) of bookings into a slice
	// 			firstNames = append(firstNames, names[0]+"|") //names[0] returns the first world of each booking and than we append it into an other slice
	// 		}
	// 		fmt.Printf("the first names of All bookings:  %v. \n", firstNames)

	// 		fmt.Printf("All bookings:  %v. \n", bookings)

	// 		if remainingTickets == 0 {
	// 			fmt.Printf("our %v is booked out. come back next year. ", conferenceName)
	// 			break
	// 		}
	// 	} else if userTickets == remainingTickets {
	// 		fmt.Printf("you are bying all the available tickets we have ! ")
	// 	} else {
	// 		fmt.Printf("we only have %v avaialble, you cant book %v. \n", remainingTickets, userTickets)
	// 	}

	// }
}


func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {

	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()

}


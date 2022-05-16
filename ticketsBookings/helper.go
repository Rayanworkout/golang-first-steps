package main

import (
	"fmt"
	"strings"
	"time"
)

func greetUsers() {
	fmt.Printf("Welcome to our %v conference !\n\n%v tickets are currently available.\n", conferenceName, remainingTickets)
}

func userInput() (string, string, uint8) {
	var userName string
	var email string
	var userTickets uint8

	fmt.Println("Entrez votre nom ...")
	fmt.Scan(&userName)

	fmt.Println("Entrez votre email ...")
	fmt.Scan(&email)

	fmt.Println("Combien de tickets ?")
	fmt.Scan(&userTickets)

	return userName, email, userTickets
}

func getFirstName() string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return strings.Join(firstNames, " / ")
}

func validateInput(name string, mail string, tickets uint8) (bool, bool, bool) {

	isValidName := len(name) >= 2 && len(mail) >= 3
	isValidEmail := strings.Contains(mail, "@")
	isValidQty := tickets <= remainingTickets && tickets > 0

	return isValidName, isValidEmail, isValidQty
}

func bookUser(name string, mail string, qty uint8) {

	var userData = UserData{
		firstName: name,
		email:     mail,
		tickets:   qty,
	}

	remainingTickets -= qty
	bookings = append(bookings, userData)

	fmt.Printf("Thanks %v, you will receive an email at %v for your %v tickets purchase.\n", name, mail, qty)

	fmt.Printf("%v tickets remaining.\n\n", remainingTickets)

	fmt.Printf("Bookings are currently: %v\n", getFirstName())
}

func sendTicket(tickets uint8, firstName string, email string) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v\n", tickets, firstName)
	fmt.Println("#################################################")
	fmt.Printf("Sending ticket: \n%v\n to %v\n", ticket, email)
	fmt.Println("#################################################")

	wg.Done()
}

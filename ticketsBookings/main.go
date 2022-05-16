package main

import (
	"fmt"
	"strings"
	"sync"
)

const conferenceName string = "Go"

var remainingTickets uint8 = 50
var bookings = make([]UserData, 0)

var wg = sync.WaitGroup{}

type UserData struct {
	firstName string
	email     string
	tickets   uint8
}

func main() {

	greetUsers()

	for remainingTickets > 0 {

		name, mail, qty := userInput()

		nameValidation, mailValidation, qtyValidation := validateInput(name, mail, qty)

		if nameValidation && mailValidation && qtyValidation {
			getFirstName()
			bookUser(name, mail, qty)

		
			wg.Add(1)
			go sendTicket(qty, name, mail)

		} else {
			invalidFields := []string{}
			if !nameValidation {
				invalidFields = append(invalidFields, "Name")
			}
			if !mailValidation {
				invalidFields = append(invalidFields, "Email")
			}
			if !qtyValidation {
				invalidFields = append(invalidFields, "Tickets quantity")
			}

			fmt.Printf("Some fields are incorrect: %v\n\nTry again.\n", strings.Join(invalidFields, " / "))
		}
	}
	wg.Wait()
}

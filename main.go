package main

import (
	"fmt"
	"go-aws/imports"
)

func main() {

	fmt.Println("hello here")
	newTicket := imports.Ticket{
		ID:    12,
		Event: "something",
	}
	fmt.Println(newTicket)
    newTicket.PrintEvent()
}

package main

import (
	"fmt"
	//s "strings"
)

func main(){
	var conferenceName = "Go Conference"
	const conferenceTicket int = 50
	var remainingTickets uint = 50
	var pf = fmt.Printf
	var pl = fmt.Println
	pf("Welcome to %v booking application\n", conferenceName)
	pf("We have a total of %v tickets and %v are still available.\n",conferenceTicket,remainingTickets)
	pl("Get your ticket to attend")

	//var userName string
	var userTickets uint
	// ask user for their name
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	fmt.Printf("You have booked %v tickets. The system has %v tickets remaining\n", userTickets,remainingTickets)
	
	
}

package main

import "fmt"

func price() int {
	return 1
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {
	switch price := price(); { 
	case price < 2:
		fmt.Println("This is a cheaply priced ticket")
	case price < 10:
		fmt.Println("This is a moderately priced ticket")
	default:
		fmt.Println("This is an expensively priced ticket")
	}

	ticket := Economy;
	switch ticket {
	case Economy:
		fmt.Println("Economy seating")
	case Business:
		fmt.Println("Business seating")
	case FirstClass:
		fmt.Println("First Class seating")
	default:
		fmt.Println("First Class seating")
	}
}

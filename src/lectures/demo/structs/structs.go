package main

import "fmt"

type Passenger struct {
	Name string
	TicketNumber int
	Boarded bool
}

type Bus struct {
	Frontseat, Backseat Passenger
}

func main() {

	Ed := Passenger{"Ed", 1, false}
	fmt.Println(Ed)

	var (
		Imy = Passenger{}
		Nick = Passenger{Name: "Nick"}
	)

	Imy.Name = "Imy"
	Imy.TicketNumber = 2
	Imy.Boarded = true

	Nick.TicketNumber = 3
	Nick.Boarded = true

	if Imy.Boarded {
		fmt.Println("Imy has boarded the bus")
	}

	if Nick.Boarded {
		fmt.Println("Imy has boarded the bus")
	}

	var Myles Passenger

	Myles.Name = "Myles"
	Myles.TicketNumber = 4

	bus := Bus {Imy, Nick}

	fmt.Println(bus)

}

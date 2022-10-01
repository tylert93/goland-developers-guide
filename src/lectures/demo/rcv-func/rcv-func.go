package main

import "fmt"

type Space struct {
	Occupied bool
}

type ParkingLot struct {
	Spaces []Space
}

func occupySpace(lot *ParkingLot, spaceNum int ) {
	lot.Spaces[spaceNum - 1].Occupied = true
}

func (lot *ParkingLot) occupySpace(spaceNum int) {
	lot.Spaces[spaceNum - 1].Occupied = true
}

func (lot *ParkingLot) vacateSpace(spaceNum int) {
	lot.Spaces[spaceNum - 1].Occupied = false
}

func main() {
	lot := ParkingLot{make([]Space, 5)}

	fmt.Println(lot)

	occupySpace(&lot, 2)
	lot.occupySpace(5)

	fmt.Println(lot)

	lot.vacateSpace(2)

	fmt.Println(lot)
}

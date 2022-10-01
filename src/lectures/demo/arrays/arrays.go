package main

import "fmt"

type Room struct {
	Name string
	IsClean bool
}

func checkForCleanRooms(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]

		if room.IsClean {
			fmt.Println(room.Name, "is clean")
		} else {
			fmt.Println(room.Name, "is dirty")
		}
	}
}

func main() {
	rooms := [...]Room{
		{Name: "Lounge"},
		{Name: "Kitchen"},
		{Name: "Bedroom"},
		{Name: "Bathroom"},
	}

	checkForCleanRooms(rooms)

	fmt.Println("Performing cleaning")
	rooms[1].IsClean = true
	rooms[3].IsClean = true

	checkForCleanRooms(rooms)
}

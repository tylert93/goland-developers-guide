package main

import "fmt"

func printSlice(title string, slice []string) {
	fmt.Println("---", title, "---")
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	fmt.Println()
}

func main() {
	route1 := []string{"Chemist", "Tailor", "Shoe shop"}

	printSlice("Original route", route1)

	route2 := append(route1, "John Lewis", "M&S")

	printSlice("Appended route", route2)

	fmt.Println(route2[0], "visited")
	fmt.Println(route2[1], "visited")
	fmt.Println()

	printSlice("Remaining route", route2[2:])
}

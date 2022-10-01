package main

import "fmt"

func main() {
	phrase := []string{"Hello", "world", "!",}

	for i, element := range phrase {
		fmt.Println(i, element, ":")

		for _, char := range element {
			fmt.Printf("%q\n", char)
		}

		fmt.Println()
	}

	
}

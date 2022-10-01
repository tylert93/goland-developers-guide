package main

import "fmt"

func main() {
	shoppingList := map[string]int{
		"eggs": 6,
		"potatoes": 4,
	}

	shoppingList["milk"] = 2
	shoppingList["bread"] += 1

	fmt.Println(shoppingList)

	fmt.Println("\nRemoving bread")
	delete(shoppingList, "bread")

	fmt.Println()
	fmt.Println(shoppingList)

	cereal, found := shoppingList["cereal"]

	if !found {
		fmt.Println("\nDon't need cereal")
	} else {
		fmt.Println("\nNeed", cereal, "cereal")
	}

	totalItems := 0

	for _, amount := range shoppingList {
		totalItems += amount
	}

	fmt.Println("\nThere are", totalItems, "items on the shopping list")
}

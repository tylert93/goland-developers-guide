//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	Name string
	Price float32
}

func getLastItem(shoppingList [4]Product) Product {
	return shoppingList[getNumberOfItems(shoppingList) - 1]
}

func getNumberOfItems(shoppingList [4]Product) int {
	total := 0

	for i := 0; i < len(shoppingList); i++ {
		item := shoppingList[i]

		if item.Name != "" {
			total += 1
		}	
	}

	return total
}

func getCostOfItems(shoppingList [4]Product) float32 {
	var sum float32 = 0.0

	for i := 0; i < len(shoppingList); i++ {
		item := shoppingList[i]
		sum += item.Price
	}

	return sum
}

func getInfo(shoppingList [4]Product) {
	fmt.Println("Last item:", getLastItem(shoppingList))
	fmt.Println("Number of items:", getNumberOfItems(shoppingList))
	fmt.Println("Total cost: £", getCostOfItems(shoppingList))
}

func main() {
	shoppingList := [4]Product{
		{"bread", 0.5},
		{"milk", 0.8},
		{"toothpaste", 2.5},
	}

	getInfo(shoppingList)

	fmt.Println("Adding item")
	shoppingList[3] = Product{"Coca cola", 1}

	getInfo(shoppingList)
}

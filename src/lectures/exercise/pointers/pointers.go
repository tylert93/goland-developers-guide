//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

type Tag struct {
	Name string
	Active bool
}

func setTag(tagName string, state bool, tags *[]Tag) {
	currentTags := *tags

	for i, currentTag := range *tags {
		if currentTag.Name == tagName {
			currentTags[i].Active = state
		}
	}
}

func checkout(tags *[]Tag) {
	currentTags := *tags

	for i := range *tags {
		currentTags[i].Active = false
	}
}

func main() {
	tags := []Tag{
		{"t-shirt", true},
		{"hoodie", true},
		{"shoes", true},
		{"beanie", true},
	}

	fmt.Println(tags)

	fmt.Println("\n----- UPDATING TAG -----")
	setTag("hoodie", false, &tags)
	fmt.Println(tags)

	fmt.Println("\n----- CHECKING OUT TAGS -----")
	checkout(&tags)
	fmt.Println(tags)
}

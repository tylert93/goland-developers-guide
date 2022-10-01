//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

const (
	CheckedIn = 0
	CheckedOut = 1
)

type BookTitle string
type Member string
type BookData struct {
	Owner Member
	Status int
	Date time.Time
}

type Library struct {
	Books map[BookTitle]BookData
	Members []Member
}

func checkoutBook(title BookTitle, member *Member, library *Library) {
	book, found := library.Books[title]

	if found {
		book.Owner = *member
		book.Status = CheckedOut
		book.Date = time.Now()

		library.Books[title] = book
	}
}

func returnBook(title BookTitle, library *Library) {
	book, found := library.Books[title]

	if found {
		book.Owner = ""
		book.Status = CheckedIn
		book.Date = time.Now()

		library.Books[title] = book
	}
	
}

func getStats(library *Library) {
	fmt.Println("------ Books ------")

	books := library.Books
	for title, data := range books {
		fmt.Println("Title:", title)

		if data.Status == CheckedIn {
			fmt.Println("Status: Checked in")
			fmt.Println("Return date:", data.Date)
		} else {
			fmt.Println("Status: Checked out")
			fmt.Println("Checkout date:", data.Date)
			fmt.Println("Checked out by:", data.Owner)
		}
		fmt.Println()
	}

	fmt.Println("------ Members ------")

	members := library.Members
	for _, member := range members {
		fmt.Println(member)
	}
	fmt.Println()
}

func main() {
	books := map[BookTitle]BookData{
		"Clean Code": {},
		"Clean Architecture": {},
		"Scopes & Closures": {},
		"Golang": {},
	}

	members := []Member{"Imy", "Ed", "Myles"}

	library := Library{books, members}

	getStats(&library)

	fmt.Println("------ Checking out books ------")
	fmt.Println()
	checkoutBook("Clean Architecture", &members[1], &library)
	checkoutBook("Golang", &members[2], &library)

	getStats(&library)

	fmt.Println("------ Returning book ------")
	fmt.Println()
	returnBook("Golang", &library)

	getStats(&library)
}

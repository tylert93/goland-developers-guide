package main

import "fmt"

func main() {
	var myName = "Tom"

	fmt.Println("My name is", myName)

	var myBrother string = "Matt"

	fmt.Println("My brother is called", myBrother)

	myEmail := "tom@gmail.com"

	fmt.Println("My email address is", myEmail)

	var myIq int

	fmt.Println("My IQ is", myIq)

	errCode, message := 404, "Page not found"

	fmt.Println("The error code is", errCode, "and the error message is", message)

	subject, message := "Resignation", "I hereby tender my resignation"

	fmt.Println("The email subject is", subject, "and the message is", message)

	var (
		myHeight = `5'11"`
		myWeight = "69Kg"
	)

	fmt.Println("My height is", myHeight, "and my weight is", myWeight)

	word1, word2, _ := "Hello", "World", "!"

	fmt.Println(word1, word2)
}

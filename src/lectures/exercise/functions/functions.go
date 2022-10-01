//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func greet(name string) {
	fmt.Println("Hello there", name)
}

func goodbye() string {
	return "au revoir missouri"
}

func add(a, b, c int) int {
	return a + b + c
}

func getNum() int {
	return 5
}

func getNums() (int, int) {
	return 11, 18
}

func main() {
	greet("Tom")

	fmt.Println(goodbye())

	num1, num2 := getNums()

	sum := add(getNum(), num1, num2)

	fmt.Println("5 + 11 + 18 =", sum)
}

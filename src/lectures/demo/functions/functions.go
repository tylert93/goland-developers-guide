package main

import "fmt"

func greet() string {
	return "Hello there"
}

func double(num int) int {
	return num + num
}

func add(lhs, rhs int) int {
	return lhs + rhs
}

func main() {
	greeting := greet()

	fmt.Println(greeting)

	dozen := double(6);

	fmt.Println("A dozen is", dozen)

	bakersDozen := add(dozen, 1)

	fmt.Println("A baker's dozen is", bakersDozen)

	anotherBakersDozen := add(double(6), 1)

	fmt.Println("Another baker's dozen is", anotherBakersDozen)
}

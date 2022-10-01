package main

import "fmt"

func main() {
	sum := 0

	for i := 0; i <= 10; i++ {
		sum += i
		fmt.Println("Increment - sum is", sum)
	}

	for sum > 10 {
		sum -= 5
		fmt.Println("Decrement - sum is", sum)
	}
}

package main

import "fmt"

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	counter.hits += 1
	fmt.Println("\nHits:", counter.hits)
}

func replace(old *string, new string, counter *Counter) {
	*old = new
	increment(counter)
}

func main() {
	greeting := "Hello"
	planet := "World"

	replacements := Counter{}

	fmt.Println(greeting, planet)

	replace(&greeting, "Hi", &replacements)

	fmt.Println(greeting, planet)

	phrase := []string{greeting, planet}

	replace(&phrase[1], "Mars", &replacements)

	fmt.Println(phrase)

	fmt.Println()
	fmt.Println(greeting, planet)
}

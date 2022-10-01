package main

import "fmt"

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {
	score1, score2, score3 := 7, 2, 9

	if score1 > score2 {
		fmt.Println("Score1 is greater then score2")
	} else if score2 > score1 {
		fmt.Println("Score2 is greater then score1")
	} else {
		fmt.Println("Score1 and score2 are equal")
	}

	if average(score1, score2, score3) > 7 {
		fmt.Println("These scores are acceptable")
	} else {
		fmt.Println("These scores are unacceptable")
	}
}

//--Summary:
//  Create a program that can perform dice rolls using various configurations
//  of number of dice, number of rolls, and number of sides on the dice.
//  The program should report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must use variables to configure:
//  - number of times to roll the dice
//  - number of dice used in the rolls
//  - number of sides of all the dice (6-sided, 10-sided, etc determined
//    with a variable). All dice must use the same variable for number
//    of sides, so they all have the same number of sides.
//
//--Notes:
//* Use packages from the standard library to complete the project
//* Try using different values for your variables to check the results

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	dice, sides, rolls := 2, 6, 3

	total := 0
	
	for i := 1; i <= rolls; i++ {

		rollSum := 0
		throws := ""

		for i := 1; i <= dice; i++ {
			throw := rand.Intn(sides -1) + 1

			rollSum += throw

			if throws == "" {
				throws += strconv.Itoa(throw)
			} else {
				throws += " / " + strconv.Itoa(throw)
			}
		}

		total += rollSum
		fmt.Println("Roll", i, "total is", rollSum, "(", throws, ")")

		switch {
		case dice == 2 && rollSum == 2:
			fmt.Println("Snake eyes");
		case rollSum == 7:
			fmt.Println("Lucky 7");
		}

		if rollSum % 2 == 0 {
			fmt.Println("Even")
			
		} else {
			fmt.Println("Odd");
		}

		fmt.Println();
	}

	fmt.Println("Overall total is", total)
	
}

//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//* Write a receiver function that performs the mathematical operation
//  on two operands
//* Operations required:
//  - Add, Subtract, Multiply, Divide
//* The existing function calls in main() represent the API and cannot be changed
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package main

import "fmt"

type Operator byte

const (
	Add Operator = iota
	Subtract
	Multiply
	Divide
)

func (o Operator) calculate(num1, num2 int) int {
	switch o {
	case Add:
		return num1 + num2
	case Subtract:
		return num1 - num2
	case Multiply:
		return num1 * num2
	case Divide:
		return num1 / num2
	default:
		return num1 + num2
	}
}

func main() {
	var (
		add Operator = Add 
		sub Operator = Subtract
		mul Operator = Multiply 
		div Operator = Divide
	)

	fmt.Println(add.calculate(2, 2)) // = 4

	fmt.Println(sub.calculate(10, 3)) // = 7

	fmt.Println(mul.calculate(3, 3)) // = 9

	fmt.Println(div.calculate(100, 2)) // = 50
}

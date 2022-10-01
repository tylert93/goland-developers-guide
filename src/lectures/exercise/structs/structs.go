//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Coordinate struct {
	X, Y int
}

type Rectangle struct {
	BottomLeft, TopRight Coordinate
}

func calcVerticalSides(rect Rectangle) int {
	return rect.TopRight.Y - rect.BottomLeft.Y
}

func calcHorizontalSides(rect Rectangle) int {
	return rect.TopRight.X - rect.BottomLeft.X
}

func calcPerimeter(rect Rectangle) int {
	return (calcHorizontalSides(rect) * 2) + (calcVerticalSides(rect) * 2)
} 

func calcArea(rect Rectangle) int {
	return calcHorizontalSides(rect) * calcVerticalSides(rect)
}

func main() {

	rectOne := Rectangle{Coordinate{0, 0}, Coordinate{4, 8}}

	fmt.Println(rectOne)
	fmt.Println("Perimeter:", calcPerimeter(rectOne))
	fmt.Println("Area:", calcArea(rectOne))

	rectTwo := Rectangle{Coordinate{0, 0}, Coordinate{8, 16}}
	fmt.Println("Perimeter:", calcPerimeter(rectTwo))
	fmt.Println("Area:", calcArea(rectTwo))
}

package main

import "fmt"

func main() {
	quantity := 4
	length, width := 1.2, 2.4
	customerName := "Brandon Lee"

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters")
}

// This program demonstrates how the above breaks if we do play around with the
// declaration process.

func mainBroken() {
	var quantity :=4 // We can only declare a variable once. Obviously, this changes based on scope.
	quantity = 4 // removing the : will make it get treated as an assignment, not a declation.
	length, width := 1.2 // We have to provide a value for every variable we assign.
	length, width = "1.2", "2.4" // Variables can only be assigned values of the same type.
	customerName := "Brandon Lee"

	// fmt.Println(customerName)
	// fmt.Println("has ordered", quantity, "sheets")
	// fmt.Println("each with an area of")
	// fmt.Println(length*width, "square meters")

	// By commenting out the above, we get a problem because all 
	// Declared variables must be used in our program. If we remove the code that uses a variable, we 
	// must also remove the declaration.
}
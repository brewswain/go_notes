package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myInteger int = 2

	fmt.Println(reflect.TypeOf(myInteger))
	fmt.Println(reflect.TypeOf(float64(myInteger)))

	fixConversionExample()
}

func fixConversionExample() {
	var length float64 = 1.2
	var width int = 2
	fmt.Println("The original length is", length)

	length = float64(width) // Trying to assign the int value to our float

	fmt.Println("After messing around with conversions, the new length is", length)
}

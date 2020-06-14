package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("42 is a type of:", reflect.TypeOf(42))
	fmt.Println("3.1415 is a type of:", reflect.TypeOf(3.1415))
	fmt.Println("true is a type of:", reflect.TypeOf(true))
	fmt.Println("\"Hello, Go!\" is a type of:", reflect.TypeOf("Hello, Go!"))
}

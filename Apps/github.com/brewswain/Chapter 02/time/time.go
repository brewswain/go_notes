package main

import (
	"fmt"
	"time" // We need to import the time package to use the time.Time type.
)

func main() {
	var now time.Time = time.Now() // time.Now returns a time.Time value of the current date and time.
	var year int = now.Year()
	fmt.Println(year)
}

package main

import (
	"fmt"
)

var a = "a"

func main() {
	a = "a"
	b := "b"
	c := "c"
	d := "d"
	if true {
		if true {
			fmt.Println(a)
			fmt.Println(b)
			fmt.Println(c)
			fmt.Println(d)
		}
		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)
		fmt.Println(d)
	}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

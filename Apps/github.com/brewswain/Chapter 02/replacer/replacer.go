package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "G# r#cks!"
	replacer := strings.NewReplacer("#", "o") // This returns a strings.Replacer that's set to replace every # with o
	fixed := replacer.Replace(broken)         // Call Replace() on the strings.Replacer, and pass it a string to do the replacements on.
	fmt.Println(fixed)
}

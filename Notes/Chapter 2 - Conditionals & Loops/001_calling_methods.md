# Calling Methods

Before we go into details of conditionals etc, we should have a basic understanding of how methods work in Go, cause they're similar enough to other languages like JS that the differences are actually more jarring.

Let's look at this example below:

```
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
```

The time.Now() function returns a new Time value for the current date and time, which we store in the `now` variable. Then, we call the Year() method on the value that `now` refers to.

## You can say therefore, that Methods are functions that are associated with values of a particular type.

Let's use another example. In the below program, we're going to want to search through a string for a substring, and replace every occurence of that substring with another string. Think of it as using regex to replace specific characters. We're going to replace every # symbol with the letter o:

```
package main

import (
    "fmt"
    "strings"
)

func main() {
    broken := "G# r#cks!"
    replacer := strings.NewReplacer("#", "o") // This returns a strings.Replacer that's set to replace every # with o
    fixed := replacer.Replace(broken) // Call Replace() on the strings.Replacer, and pass it a string to do the replacements on.
    fmt.Println(fixed)
}
```

In this case, the strings.NewReplacer function takes arguments with a string to replace("#"), and a string to replace it with("o"), and returns a strings.Replacer. When we pass a string the that Replacer value's Replace method, it returns a string with those replacements made.

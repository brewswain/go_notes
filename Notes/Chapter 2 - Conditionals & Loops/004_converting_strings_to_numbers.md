# Converting Strings to numbers

Returning to our grading program logic that we've been writing, let's add some logic that'll determine whether the grade that the user inputs is passing or failing:

```
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Enter a grade:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	if input >= 60 {
        status := "passing"
    } else {
        status := "failing"
    }
}
```

In this current form, we'll get a compilation error:

```
.\pass_fail.go:19:11: cannot use 60 (type untyped int) as type string
```

The problem's laid out pretty clearly for us: Input from the keyboard is read in as a string, when we obviously need it to be read as a number. We also don't have a direct path to convert from string to number, so we can't do this:

```
float64("2.6")
```

The above would fail quite spectacularly, so how do we do this? Well, we begin by isolating the issues here. Firstly, our input string still has '\n' attached to it from when the user pressed the enter Key. Then, we need to convert the remainder of the string to a float.

Thankfully, we can solve the first issue by using a function called TrimSpace() that we get from the `strings` package. As the name sounds, it'll remove all whitespace characters from the start and end of a string.

That's pretty easy, we'll just pass `input` through TrimSpace() and reassign it back to the input variable like so:

```
input = string.TrimSpace(input)
```

Easy-Peasy! Now, let's get to the main problem at hand, the string conversion. To accomplish this, we're going to use a package called `strconv`, which has the function ParseFloat() that'll do the work fur us:

```
grade, err := strconv.ParseFloat(input, 64)
```

Please note the (input, 64). input refers to the string we wish to convert, and the 64 refers to the number of bits of precision we want for the result. It should be noted that 64 will generally be the preferred option here.

You probably also noted the `err` return value, which will be nil unless there was a problem converting the string, so pretty standard fare.

Let's update our program!

```
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}

	if grade >= 60 {
		status := "passing"
	} else {
		status := "failing"
	}
}

```

Nice! Let's tackle the remainder of our program in the next file.

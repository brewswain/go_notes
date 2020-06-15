# Blocks

Now that we've converted the user's grade input over to a float64 value, the conditional should be able to determine if it's passing or failing.

However, we have a pretty obvious error that we get:

```
.\pass_fail.go:28:3: status declared but not used
.\pass_fail.go:30:3: status declared but not used
```

Pretty clear error, we haven't used the `status` variables yet. Interestingly, we get the error twice, but let's ignore that for now.

Let's fix this problem by printing the percentage grade that we were given, as well as the status.

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
    fmt.Println("A grade of", grade, "is", status)
}

```

However, now we get a new error!

```
.\pass_fail.go:32:41: undefined: status
```

Well, I assume that you know about scope so I won't really go over it, especially since this issue is pretty obvious. This is also why we got the error code twice. Within a broad term for Go, code blocks are usually surrounded by curly braces {}, although there are also blocks at the source code file and package levels.

A good solution seems to be to declare our status above our conditional blocks (treat it like manually hoisting):

```
    var status string

	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
    fmt.Println("A grade of", grade, "is", status)
```

The only thing to keep in mind here is that we made sure to change our variable declarations (:=) to assignment (=) inside of ouor conditionals.

And now, at long last, our program should work!

# Let's make a game!

Our goal is to make a variant of "High/Low" with the number pool being from 1 - 100.

---

## Random Number Generation

Let's begin by importing a package to randomly generate said number:

```
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	target := rand.Intn(100) + 1
	fmt.Println(target)
}

```

Pretty simple so far, we're using a more specific import path to get our `rand` package. rand.Intnt() will return a random integer between 0 and the number we've provided, hence why we add 1 to the value to shift it from returning 0-99 to 1-100.

Now, if you were like me and instantly compiled and ran this program, you would notice that although it works and generates us a random number, it generates the exact same number over and over, which kind of defeats the purpose.

To get different random numbers, we're going to use a new function, Seed(). This will give our RNG a value that it'll use to generate other random values. However, if we keep giving it the same seed value then it'll keep giving us the same random values, which means that we wouldn't have actually accomplished anything.

To bypass that, we can actually finesse the `time` package to give us a constantly changing package, like so:

```
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1

	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	fmt.Println(target)

}
```

Nice! this should give us a new number each time. For reference by the way, the Unix() method on Time() actually converts our time to the Unix time format, which is an integer with the number of seconds since Jan. 1st, 1970. Pretty cool actually!

We then pass that dynamic integer to our Seed() and hey presto!

---

## Get integer from keyboard

Much like how we got an integer from our grading program (pass_fail.go), we'll want to use `bufio1`, `log`, `os`, `strconv`, and `strings` to handle both our input and our conversion.

Obviously, in this case we don't want to convert our string into a float, but an integer, so keep that in mind:

```
package main

import (
    "bufio"
	"fmt"
    "log"
	"math/rand"
    "os"
    "strconv"
    "strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	fmt.Println(target)

    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Make a guess: ")
    input, err := reader.ReadString('\n')
    if err != nil {
        log.Fatal(err)
    }
    input = strings.TrimSpace(input)
    guess, err := strconv.Atoi(input)
    if err != nil {
        log.Fatal(err)
    }
}
```

---

## Compare our guess to the target

So now, we just want to compare the user's guess to the tandomly generated number, and tell them whether it was higher or lower. This should be a cinch with conditionals:

```
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	fmt.Println(target)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Make a guess: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	guess, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}

	if guess < target {
		fmt.Println("Sorry, your guess was low.")
	} else if guess > target {
		fmt.Println("Sorry, your guess was High.")
	} else {
		fmt.Println("You got it!!!")
	}
}

```

---

## Loops

We currently only let the player guess once, but let's let them guess up to 10 times. For this, we're going to use a loop. A for loop, to be specific. Actually, on that note, please be aware that Golang only uses for loops!

I'm gonna assume that we already know how for loops work, so i'm just gonna write up a dummy one for example's sake:

```
for x := 1; x <= 6; x++ {
    fmt.Println("x is now", x)
}
```

the really interesting thing that go offers with loops are the assignment operators `+=` and `-=`, They take the value in a variable, add or subtract another value, then assign the result back to the variable.

This allows us to increment a loop in amounts other than 1.

```
for x :=1; x <=5; x += 2 (
    fmt.Println(x)
)
```

Let's use loops in our game:

```
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	fmt.Println(target)

	reader := bufio.NewReader(os.Stdin)

    for guesses := 0; guesses < 10; guesses++ {
        fmt.Println("You have", 10-guesses, "guesses left.")


        fmt.Print("Make a guess: ")
        input, err := reader.ReadString('\n')
        if err != nil {
            log.Fatal(err)
        }
        input = strings.TrimSpace(input)
        guess, err := strconv.Atoi(input)
        if err != nil {
            log.Fatal(err)
        }

        if guess < target {
            fmt.Println("Sorry, your guess was low.")
        } else if guess > target {
            fmt.Println("Sorry, your guess was High.")
        } else {
            fmt.Println("You got it!!!")
        }
    }
}

```

Now our program prompts for a guess 10 times, then it ends! Of course, we still have a problem here, where even if you pick the correct choice, it still loops 10 times.

Let's fix that with `continue` and `break` statements.

### Continue

`continue` immediately skips to the next iteration of a loop without running any further code in the loop's block:

```
for x := 1; x <=3; x++ {
    fmt.Println("before continue")
    continue
    fmt.Println("after continue")
}
```

The above will print "before continue" 3 times since the `continue` statement will skip back to the top of the loop.

### Break

`break` immediately breaks out of a loop. No further code in the loop is executed, and no iterations are run. Execution moves to the first statement following the loop:

```
for x := 1; x <= 3; x++ {
    fmt.Println("before break")
    break
    fmt.Println("after break")
}
fmt.Println("after loop")
```

In this case, our loop will print "before break" once, then break out of the loop and print "after loop".

Let's do this in our game! The application is pretty simple:

```
// snip

 if guess < target {
            fmt.Println("Sorry, your guess was low.")
        } else if guess > target {
            fmt.Println("Sorry, your guess was High.")
        } else {
            fmt.Println("You got it!!!")
            break
        }

```

Now, when the player guesses correctly, the loop will exit.

---

## Revealing the target

We're nearly done! Currently, if the player makes 10 guesses without getting the target, the loop will exit. We want to make the program print a message in that case that'll tell them that they lost, then tell them what the target was.

However, if the player wins, the loop also exits. We obviously don't want to tell the player that they've lost when they've won,
so let's set up a flag that'll tell our program if we've won or lost:

```
// snip

    success := false // We declare this outside the loop to ensure that we avoid any scope issues

    for guesses := 0; guesses < 10; guesses++ {
        fmt.Println("You have", 10-guesses, "guesses left.")


        fmt.Print("Make a guess: ")
        input, err := reader.ReadString('\n')
        if err != nil {
            log.Fatal(err)
        }
        input = strings.TrimSpace(input)
        guess, err := strconv.Atoi(input)
        if err != nil {
            log.Fatal(err)
        }

        if guess < target {
            fmt.Println("Sorry, your guess was low.")
        } else if guess > target {
            fmt.Println("Sorry, your guess was High.")
        } else {
            success = true
            fmt.Println("You got it!!!")
            break
        }
    }

    if !success {
        fmt.Println("Sorry, you didn't guess the correct number. It was:", target)
    }
```

And thus we've very simply set a conditional failure message that renders if success = false. Let's fix the code for production by removing the line that prints what the actual target is, and we should be done:

```
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")

	reader := bufio.NewReader(os.Stdin)

	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")

		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Sorry, your guess was low.")
		} else if guess > target {
			fmt.Println("Sorry, your guess was High.")
		} else {
			success = true
			fmt.Println("Congrats, you got it!!!")
			break
		}
	}

	if !success {
		fmt.Println("Sorry, you didn't guess the correct number. It was:", target)
	}
}

```

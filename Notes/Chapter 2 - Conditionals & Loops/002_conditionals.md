# Conditionally running code

I assume you already know why you'd want conditionally run code, so let's just look at the scenario we're going to use to demonstrate this.

We need to write a program that allows a student to type in their percentage grade and tells them whether they passed or not. Pretty simple concept, we'll put different responses based on whether or not the user passes or fails.

Let's look at the code that we'll start with:

```
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    fmt.Print("Enter a grade:")
    reader := bufio.NewReader(os.Stdin)
    input := reader.ReadString('\n')
    fmt.Println(input)
}
```

Now, don't do like me and instantly try to run the above code, because it WILL return an error. More on that later.

You may have noticed that we used fmt.Print() and not Println(). This is because Print() doesn't skip to a new terminal line after printing a message, which is optimal for user input.

Our next line is kinda funky, but know that for **now**, we don't really need to learn the specifics of how bufio works(I hate saying that but...the show must go on), but here just know that we're storing a bufio.Reader inside of our `reader` variable. The `os.Stdin` will allow the Reader to read from standard input, which in this casae is the keyboard.

Our input := reader.Readstring('\n') will return what the user typed as a **string**. Everything up until the newline rune will be read. Using that rune to mark the end of the input is actually how the ReadString() method works, as we want to read everything that the user types up until tjhey press Enter.

## But.

As we noted, trying to run this code in its current state will actually give us an error:

```
.\pass_fail.go:12:8: assignment mismatch: 1 variable but reader.ReadString returns 2 values
```

We're trying to read the user's keyboard input, but we get an error. The problem is that the ReadString method is trying to return two values, and we've only provided one variable to assign a value to.

This is interesting to note, because this suggests that in Go, methods can return any number of values. Most commonly in Go, we use this to return an error value that can be consulted to find out if anything went wrong while the function or method was running.

A couple examples would be:

```
bool, err := strconv.ParseBool("true") // Returns an error if the string can't be converted to boolean

file, err := os.Open("myfile.txt) // Returns an error if the file can't be opened

response, err := http.Get("https://golang.org") // Returns an error if the page can't be retrieved.
```

However, remember something that we've noted before. Go requires that **every** variable that gets declared must also get used somewhere in the program. If we add an `err` variable and then don't check it, our code won't compile.

This is actually a good thing as this actually helps us detect and fix bugs!
Now, as far as how we implement this inside of our program, we have 2 main options. Ignore the error return with a blank identifier, or handle the error. Obviously, the latter is the better choice but let's go over both approaches as we mentioned something interesting, a `blank identifier`.

---

## Option 1: Use blank identifier

The ReadString method returns a second value along with the user's input, and we need to do something with that second value. We've tried just adding a second variable and ignoring it, but our code still won't compile:

```
input, err := reader.ReadString('\n') // Error: err declared and not used
```

When we have a value that would normally be assigned to a variable, but we don't intend to use it, we can use Go's **blank identifier**. Asigning a value to the blank identifier essentially discards it(while making it obvious to others reading your code that you're doing so).

To use the blank identifier, we only need to type a single underscore(\_) character in an assignment statement where we would normally type a variable name. Let's try using the blank identifier in our code:

```
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    fmt.Print("Enter a grade:")
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    fmt.Println(input)
}
```

If we try this change out and run the program, it'll work! However, ignoring an error...obviously is kind of sloppy. We usually would like our program to let us know if errors occur, or at least I would, so let's look at error handling.

## Option 2: Handle the error

If we get an error back from the ReadString method, the blank identifier would just cause the error to be ignored, and our program would proceed anyway, with some potentially invalid data.
In this case, it would be more appropriate to alert the user and stop the program if there was an error.

Fortunately, there's a package called `log` that can do both of these operations for us! `log` has a Fatal() function that logs a message to the terminal _and_ stops the program. Let's remove our blank identifier and replace it with our `err` variable so that we're recording the error again. This time however, let's use the Fatal() function to log the error and half the program:

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
    log.Fatal(err) // Reports the error and stops the program
    fmt.Println(input)
}
```

The program itself runs, **but!** We have a new problem, which leads us back to the crux of this entire lesson(We're learning about conditionals, remember?)

As it stands, We want our program to report an error and stop running if we encounter a problem reading input from the keyboard. However, it currently stops running even when everything's working correctly 😔

You may have noticed that the error message we return is an extremely loquacious... `nil`.
Yeah. That's our error message.

This is due to functions and methods like Readstring() returning an error value of nil to suggest that there's nothing there. So basically, `err` being `nil` suggests that there's no error. However, our program is currently setup to just report the `nil` error. We actually want the program to exit if `err` has a value _other_ than `nil`.

We can do this by using (Ta-da!) Conditionals. Again, the assumption is that we understand how conditionals work, so let's look at this example if block:

```
if 1 < 2 {
    fmt.Println("It's true!")
}
```

Like most other languages, we can use `if else` blocks:

```
if grade == 100 {
    fmt.Println("Perfect!")
} else if grade >= 60 {
    fmt.Println("You passed!")
} else {
    fmt.Println("L")
}
```

Of course, just like javascript, we can do some cool stuff to make sure that code executes when false, by using the bang operator:

```
if !true {
    fmt.Println("beep")
}
```

We can also use && and || for our AND + OR operators. Pretty sweet!

Now, something you probably noticed was that our statements don't use parantheses( unless we're using them to set an order of operations). That's pretty interesting to note.

## Logging our fatal error

Back to our issue at hand, let's use our `if` statements to log an error and exit only if `err` isn't `nil`.

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
    fmt.Println(input)
}
```

If we run our program now, we'll see that it's working again. Also, if we have any errors reading user input, we'll see those as well!

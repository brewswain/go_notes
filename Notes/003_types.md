# Types

Values in Go are stored in different Types, which specify what the values can be used for.
Integers can be used for math operations, but strings can't, etc.

Go is **Statically Typed**, which means that it knows what the
types of your values are even before your program runs. Therefore, functions expect their arguments to be of particular types, and their return values have types as well (Which may **or** may not be the same as the argument types). This means that Go will give you an error message if you use the wrong type of value in the wrong place (practise for learning Typescript!)

We can use the reflect package's TypeOf() function to view the type of any value:

```
package main

import (
    "fmt"
    "reflect"
)

func main() {
    fmt.Println(reflect.TypeOf(42))
}
```

# Shadowing Names

If you're like me (considering that I wrote these notes for myself, that chance is pretty high), the fact that we're using `err` instead of `error` as our error variable is highly annoying as abbreviations are (to me) often simply lazy naming that makes it a bit harder to read code.

### However.

In this case, naming a variable `error` would actually be an awful idea, as it would shadow the name of a _type_ called `error`.

When we declare a variable, we should make sure that it doesn't share a name with any existing functions, packages, ttypes, or other variables. Obviously, as this is my first foray into a typed language, This wasn't even a consideration for me and now I understand why people tend to use the `err` convention, even in languages that don't require it.

The reason why _shadowing_ is bad is that if something by the same name exists in the enclosing scope, our variable will shadow it, taking precedence over it.

Check out the example of shadowing messing stuff up for us below:

```
package main

import "fmt"

func main() {
    var int int = 12
    var append string = "minutes of bonus footage"
    var fmt string = "DVD"
}
```

By themselves, they don't cause any errors. However, once we try to access any type, function or package that the variables are shadowing, we'll get the value in the variable instead which results in some delicious compilation errors:

```
package main

func main() {
	var int int = 12
	var append string = "minutes of bonus footage"
	var fmt string = "DVD"

	var count int
	var languages = append([]string{}, "English")
	fmt.Println(int, append, "on", fmt, languages)
}

```

```
.\shadowing_errors.go:8:6: int is not a type
.\shadowing_errors.go:9:24: cannot call non-function append (type string), declared at .\shadowing_errors.go:5:6
.\shadowing_errors.go:10:5: fmt.Println undefined (type string has no field or method Println)
```

If this doesn't show you why shadowing sucks then you should probably just go wherever your heart tells you to and live freely tbh.

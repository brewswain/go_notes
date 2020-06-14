# Variables

quite simply, you declare variables in Go using the `var` kayword.

Now, the interesting thingm about Go is that we can also declare the type of values that the variable will hold:

```
var quantity int

var length, width float64

var customerName string
```

As you can see with "var length, width float64", we can actually declare multiple variables of the same type as once!

This rule also applies with assigning values to multiple variables, like so:

```
var length, width float64

length, width = 1.2, 2.4
```

Of course, as with standard variable assignation conventions, we can place the value inline:

```
var length, width float64 = 1.2, 2.4
```

It shoulpd be noted that if we assign a value to a variable at the same time that w declare it, we can usually omit the variable type from the declaration, as the type of the value assigned will be used as the type of that variable:

```
var quantity = 4
var length, width = 1.2, 2.4
var customerName = "Damon"
```

---

## Zero Values

Finally, if we declare a variable without assigning it a value, the variable will contain the zero value for its type. For instance, integers would be 0, strings would be "", booleans would be false, etc.

---

## Short Variable Declarations

As we discussed earlier, we can declare variables and assign values to them on the same line:

```
var quantity int = 4
```

However. If we know what the initial value of a variable is going to be as soon as we declare it, it's more typical to use a **short variable declaration**.

Instead of explicitly declaring the type of the variable and assigning to it later with `=`, we can do both at once by using `:=`.

Let's update our previous examples to show what we mean:

```
quantity := 4
length, width := 1.2, 2.4
customerName := "Damon"
```

There's now no need to explicitly declare the variable's type; **the type of the value assigned to the variable becomes the type of that variable.**

Because short variable declarations are so convenient and concise, we'll probably see them used more often then regular declarations.

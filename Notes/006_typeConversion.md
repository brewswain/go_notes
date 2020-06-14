# Type Conversions

Just like other languages, math and comparison operations in Go require that the included values be of the same type. If they're not you'll get an error message:

```
invalid operation: foobar (mismatched types)
```

The same is true of assigning new values to variables:

```
var length float64 = 1.2
var width int = 2
length = width  // Trying to assign the int value to our float

fmt.Println(length)
```

The above will produce an error!:

```
cannot use width (type int) as type float64 in assignment
```

The solution is to use **conversions**, which let you convbert a value from one type to another by simply providing the type to you want to convert a value to, followed by the value itself in parantheses:

```
var integer int = 2

float64(integer)
```

The result is a new value of the desired type.
Look at Conversions/ in our apps folder for an example of this.

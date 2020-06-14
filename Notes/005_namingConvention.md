# Naming Convention

use camelCase and always start your name with a letter.

Other than that, if the name of a variable, function or type begins with a capital letter, it's considered `exported` amd can be accessed from packages outside the current one:

```
fmt.Println()
```

The P in fmt.Println() is capitalized so that it can be used from the main package or any other.

If a variable/.function begins with a lowercase letter, it's considered unexported and can only be accessed within the current package.
